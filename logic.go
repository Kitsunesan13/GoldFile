package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FileItem struct {
	Entry   os.DirEntry
	RelPath string
}

type StateAplikasi struct {
	CurrentDir     string
	Files          []os.DirEntry 
	
	FilteredFiles  []FileItem    
	
	FileIndex      int 
	DashboardIdx   int 
	SettingsIdx    int 
	Mode           int 
	LastOutput     string 
	
	IsEditing      bool      
	EditCursorPos  int       
	
	IsSearching    bool      
	SearchQuery    string    
	
	AsciiCache     string 
	LastImageUsed  string 
}

func InitState() StateAplikasi {
	wd, _ := os.Getwd()
	s := StateAplikasi{
		CurrentDir:   wd,
		Mode:         0, 
		LastOutput:   "Siap. Gunakan koma (,) untuk filter ganda.",
	}
	s.LoadFiles()
	return s
}

func (s *StateAplikasi) LoadFiles() {
	entries, err := os.ReadDir(s.CurrentDir)
	if err != nil {
		s.LastOutput = "Error: " + err.Error()
		return
	}
	s.Files = entries
	
	s.IsSearching = false
	s.SearchQuery = ""
	
	s.FilteredFiles = make([]FileItem, 0, len(entries))
	for _, e := range entries {
		s.FilteredFiles = append(s.FilteredFiles, FileItem{
			Entry:   e,
			RelPath: e.Name(), 
		})
	}
	
	s.FileIndex = 0
}

func (s *StateAplikasi) PerformSearch() {
	if s.SearchQuery == "" {
		s.FilteredFiles = make([]FileItem, 0, len(s.Files))
		for _, e := range s.Files {
			s.FilteredFiles = append(s.FilteredFiles, FileItem{Entry: e, RelPath: e.Name()})
		}
		s.FileIndex = 0
		return
	}

	var results []FileItem
	
	rawKeywords := strings.Split(strings.ToLower(s.SearchQuery), ",")
	var keywords []string
	for _, k := range rawKeywords {
		trimmed := strings.TrimSpace(k)
		if trimmed != "" { keywords = append(keywords, trimmed) }
	}
	
	if len(keywords) == 0 { return }

	maxResults := 100        
	maxDepth := 5            
	baseDirDepth := strings.Count(s.CurrentDir, string(os.PathSeparator))

	filepath.WalkDir(s.CurrentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil { return nil }
		
		if len(results) >= maxResults {
			return filepath.SkipDir 
		}

		currDepth := strings.Count(path, string(os.PathSeparator)) - baseDirDepth
		if currDepth > maxDepth {
			if d.IsDir() { return fs.SkipDir }
			return nil
		}

		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" || name == ".config" || name == "dist" || name == "build" || name == "tmp" {
				return filepath.SkipDir
			}
			return nil 
		}

		relPath, _ := filepath.Rel(s.CurrentDir, path)
		
		isTotalMatch := true 

		for _, key := range keywords {
			isKeywordFound := false

			if strings.Contains(strings.ToLower(relPath), key) {
				isKeywordFound = true
			}

			if !isKeywordFound {
				info, _ := d.Info()
				if info.Size() < 500*1024 { 
					content, err := readFileLimit(path, 10*1024)
					if err == nil {
						if strings.Contains(strings.ToLower(string(content)), key) {
							isKeywordFound = true
						}
					}
				}
			}

			if !isKeywordFound {
				isTotalMatch = false
				break 
			}
		}

		if isTotalMatch {
			results = append(results, FileItem{
				Entry:   d,
				RelPath: relPath,
			})
		}

		return nil
	})
	
	s.FilteredFiles = results
	s.FileIndex = 0 
}

func readFileLimit(path string, limit int64) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil { return nil, err }
	defer f.Close()
	
	stat, _ := f.Stat()
	if stat.Size() == 0 { return nil, nil }

	readSize := limit
	if stat.Size() < limit { readSize = stat.Size() }

	buf := make([]byte, readSize)
	n, err := f.Read(buf)
	if err != nil && n == 0 { return nil, err }
	
	return buf[:n], nil
}


func (s *StateAplikasi) HandleInputDashboard(key string) bool {
	if len(key) > 0 && key[0] == 3 { return false }
	if key == "\033[A" || key == "k" { if s.DashboardIdx > 0 { s.DashboardIdx-- } } else 
	if key == "\033[B" || key == "j" { if s.DashboardIdx < 3 { s.DashboardIdx++ } } else 
	if key == "\r" || key == "\n" { 
		switch s.DashboardIdx {
		case 0: s.Mode = 1 
		case 1: s.Mode = 2 
		case 2: s.Mode = 3 
		case 3: return false 
		}
	}
	return true
}

func (s *StateAplikasi) HandleInputFileManager(key string, cfg *KonfigurasiApp) {
	if s.IsSearching {
		if key == "\r" || key == "\n" || key == "\033" { s.IsSearching = false; return }
		
		if key == "\u007f" || key == "\b" {
			if len(s.SearchQuery) > 0 {
				s.SearchQuery = s.SearchQuery[:len(s.SearchQuery)-1]
				s.PerformSearch() 
			}
			return
		}

		if len(key) == 1 && key[0] >= 32 {
			s.SearchQuery += key
			s.PerformSearch() 
		}
		return
	}

	
	if key == "/" { s.IsSearching = true; s.SearchQuery = ""; s.PerformSearch(); return }
	if key == "q" { s.Mode = 0; s.LastOutput = "Kembali ke Dashboard."; return }
	
	if key == "\033[A" || key == "k" { 
		if s.FileIndex > 0 { s.FileIndex-- }
		return 
	} 
	
	if key == "\033[B" || key == "j" { 
		if s.FileIndex < len(s.FilteredFiles)-1 { s.FileIndex++ }
		return 
	} 
	
	if key == "\r" || key == "\n" || key == "\033[C" || key == "l" { s.BukaFileAtauFolder(cfg); return } 
	
	if key == "\u007f" || key == "\b" || key == "\033[D" || key == "h" { 
		prevDir := filepath.Dir(s.CurrentDir)
		if prevDir != s.CurrentDir { s.CurrentDir = prevDir; s.FileIndex = 0; s.LoadFiles() }
		return
	} 
	if key == "\t" { s.JalankanPreview() } 
}

func (s *StateAplikasi) HandleInputSettings(key string, cfg *KonfigurasiApp) {
	if s.IsEditing {
		if key == "\r" || key == "\n" { s.IsEditing = false; SaveConfig(cfg); s.LastImageUsed = ""; return }
		var currentText string
		if s.SettingsIdx == 1 { currentText = cfg.ImagePath } else 
		if s.SettingsIdx == 2 { currentText = cfg.FooterText } else 
		if s.SettingsIdx > 2 { currentText = cfg.Dialogues[s.SettingsIdx-3] }
		runes := []rune(currentText)
		if key == "\033[D" { if s.EditCursorPos > 0 { s.EditCursorPos-- }; return }
		if key == "\033[C" { if s.EditCursorPos < len(runes) { s.EditCursorPos++ }; return }
		if key == "\u007f" || key == "\b" {
			if s.EditCursorPos > 0 && len(runes) > 0 {
				runes = append(runes[:s.EditCursorPos-1], runes[s.EditCursorPos:]...)
				s.EditCursorPos--
				updateConfigString(s, cfg, string(runes))
			}
			return
		}
		if len(key) == 1 && key[0] >= 32 {
			runes = append(runes, 0)
			copy(runes[s.EditCursorPos+1:], runes[s.EditCursorPos:])
			runes[s.EditCursorPos] = rune(key[0])
			s.EditCursorPos++
			updateConfigString(s, cfg, string(runes))
		}
		return
	}
	if key == "q" || key == "b" { s.Mode = 0 } else 
	if key == "\033[A" || key == "k" { if s.SettingsIdx > 0 { s.SettingsIdx-- } } else 
	if key == "\033[B" || key == "j" { if s.SettingsIdx < 2 + len(cfg.Dialogues) { s.SettingsIdx++ } } else 
	if key == "\r" || key == "\n" { 
		if s.SettingsIdx == 0 {
			switch cfg.EditorDefault {
			case "nvim": cfg.EditorDefault = "helix"
			case "helix": cfg.EditorDefault = "vim"
			case "vim": cfg.EditorDefault = "nano"
			default: cfg.EditorDefault = "nvim"
			}
			SaveConfig(cfg) 
		} else { s.IsEditing = true; var txt string; if s.SettingsIdx == 1 { txt = cfg.ImagePath } else if s.SettingsIdx == 2 { txt = cfg.FooterText } else if s.SettingsIdx > 2 { txt = cfg.Dialogues[s.SettingsIdx-3] }; s.EditCursorPos = len([]rune(txt)) }
	}
}

func updateConfigString(s *StateAplikasi, cfg *KonfigurasiApp, newVal string) {
	if s.SettingsIdx == 1 { cfg.ImagePath = newVal } else 
	if s.SettingsIdx == 2 { cfg.FooterText = newVal } else 
	if s.SettingsIdx > 2 { cfg.Dialogues[s.SettingsIdx-3] = newVal }
}

func (s *StateAplikasi) HandleInputAbout(key string) { s.Mode = 0 }

func (s *StateAplikasi) BukaFileAtauFolder(cfg *KonfigurasiApp) {
	if s.FileIndex >= len(s.FilteredFiles) { return }
	
	selected := s.FilteredFiles[s.FileIndex] 
	fullPath := filepath.Join(s.CurrentDir, selected.RelPath) 
	
	if selected.Entry.IsDir() { 
		s.CurrentDir = fullPath; s.FileIndex = 0; s.LoadFiles() 
	} else { 
		s.BukaEditor(cfg.EditorDefault, fullPath) 
	}
}

func (s *StateAplikasi) BukaEditor(editor string, path string) {
	cmdName := editor
	if editor == "helix" { if _, err := exec.LookPath("hx"); err == nil { cmdName = "hx" } else { cmdName = "helix" } }
	DisableRawMode() 
	cmd := exec.Command(cmdName, path)
	cmd.Stdin = os.Stdin; cmd.Stdout = os.Stdout; cmd.Stderr = os.Stderr
	ClearScreen()
	cmd.Run()
	EnableRawMode()
	fmt.Print("\033[2J") 
}

func (s *StateAplikasi) JalankanPreview() {
	if s.FileIndex >= len(s.FilteredFiles) { return }
	
	fileItem := s.FilteredFiles[s.FileIndex]
	fullPath := filepath.Join(s.CurrentDir, fileItem.RelPath)
	
	if !fileItem.Entry.IsDir() {
		ext := strings.ToLower(filepath.Ext(fileItem.RelPath))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" || ext == ".gif" {
			cmd := exec.Command("chafa", fullPath, "--format=symbols", "--size=40x20")
			out, err := cmd.Output()
			if err != nil { s.LastOutput = "Error image: " + err.Error() } else { s.LastOutput = string(out) }
		} else {
			content, err := os.ReadFile(fullPath)
			if err != nil { s.LastOutput = "Error read: " + err.Error(); return }
			strContent := string(content)
			if len(strContent) > 3000 { strContent = strContent[:3000] + "\n... (Large File)" }
			var highlighted strings.Builder
			lines := strings.Split(strContent, "\n")
			for _, line := range lines { highlighted.WriteString(HighlightCode(line)); highlighted.WriteString("\n") }
			s.LastOutput = highlighted.String()
		}
	} else { s.LastOutput = fmt.Sprintf("\033[34m[FOLDER]\n%s\033[0m", fileItem.RelPath) }
}