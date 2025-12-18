package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func truncateString(s string, maxLen int) string {
	if maxLen <= 0 { return "" }
	s = strings.ReplaceAll(s, "\r", "")
	if utf8.RuneCountInString(s) > maxLen {
		runes := []rune(s)
		return string(runes[:maxLen])
	}
	return s
}

func RenderEditableText(text string, cursorPos int, isFocused bool) string {
	if !isFocused { return fmt.Sprintf("\"%s\"", text) }
	var sb strings.Builder
	sb.WriteString("\"") 
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if i == cursorPos { sb.WriteString("\033[7m"); sb.WriteRune(runes[i]); sb.WriteString("\033[0m") } else { sb.WriteRune(runes[i]) }
	}
	if cursorPos == len(runes) { sb.WriteString("\033[7m \033[0m") }
	sb.WriteString("\"") 
	return sb.String()
}

func RenderManager(width, height int, frame int, state *StateAplikasi, cfg KonfigurasiApp) {
	switch state.Mode {
	case 0: RenderDashboard(width, height, frame, state, cfg)
	case 1: RenderFileManagerUI(width, height, state, cfg)
	case 2: RenderSettingsUI(width, height, state, cfg)
	case 3: RenderAboutUI(width, height, cfg)
	}
}

func RenderFileManagerUI(width, height int, state *StateAplikasi, cfg KonfigurasiApp) {
	var sb strings.Builder
	sb.WriteString("\033[H\033[2J") 
	sb.WriteString(cfg.Tema.Background)
	mainHeight := height - 5 

	sb.WriteString(MoveCursorStr(1, 1))
	menuBar := fmt.Sprintf(" GOLD v2 | %s | [Tab] Preview", cfg.EditorDefault)
	sb.WriteString(fmt.Sprintf("%s\033[30m%-*s", cfg.Tema.StatusBarBg, width, menuBar))

	sb.WriteString(MoveCursorStr(1, 2))
	sb.WriteString(fmt.Sprintf("%s%s Path: %-*s", cfg.Tema.Background, cfg.Tema.Directory, width-8, state.CurrentDir))

	midX := width / 2
	startIdx := 0
	if state.FileIndex >= mainHeight { startIdx = state.FileIndex - mainHeight + 1 }
	leftPanelWidth := midX - 4 

	fileList := state.FilteredFiles

	for i := 0; i < mainHeight; i++ {
		idx := startIdx + i
		yPos := 4 + i
		sb.WriteString(MoveCursorStr(2, yPos))
		sb.WriteString(strings.Repeat(" ", midX-2)) 
		sb.WriteString(MoveCursorStr(2, yPos))

		if idx < len(fileList) {
			fileItem := fileList[idx]
			name := TruncateAnsi(fileItem.RelPath, leftPanelWidth)
			
			icon := "  "
			if len(state.SearchQuery) > 0 { icon = "🔍" }

			if idx == state.FileIndex {
				sb.WriteString(fmt.Sprintf("%s%s ➤ %s %s", cfg.Tema.SelectedBg, cfg.Tema.SelectedFg, name, "\033[0m"))
			} else {
				col := cfg.Tema.File
				if fileItem.Entry.IsDir() { col = cfg.Tema.Directory }
				sb.WriteString(fmt.Sprintf("%s%s %s%s", cfg.Tema.Background, col, icon, name))
			}
		}
	}

	for y := 3; y < 4+mainHeight; y++ { 
		sb.WriteString(MoveCursorStr(midX, y))
		sb.WriteString(fmt.Sprintf("%s│", cfg.Tema.Border)) 
	}

	sb.WriteString(MoveCursorStr(midX+2, 3))
	sb.WriteString(fmt.Sprintf("%sPREVIEW / OUTPUT:", cfg.Tema.Directory))
	
	lines := strings.Split(state.LastOutput, "\n")
	previewMaxWidth := width - midX - 3 
	for i := 0; i < mainHeight; i++ {
		yPos := 4 + i
		sb.WriteString(MoveCursorStr(midX+2, yPos))
		sb.WriteString(strings.Repeat(" ", previewMaxWidth+1)) 
		sb.WriteString(MoveCursorStr(midX+2, yPos))
		if i < len(lines) {
			safeLine := TruncateAnsi(lines[i], previewMaxWidth)
			sb.WriteString(fmt.Sprintf("%s%s", cfg.Tema.Foreground, safeLine))
		}
	}

	searchY := height - 1
	sb.WriteString(MoveCursorStr(1, searchY))
	sb.WriteString(strings.Repeat(" ", width)) 
	sb.WriteString(MoveCursorStr(1, searchY))

	if state.IsSearching {
		sb.WriteString(fmt.Sprintf("\033[43;30m SEARCH: %s\033[5m_\033[0m", state.SearchQuery)) 
	} else if len(state.SearchQuery) > 0 {
		sb.WriteString(fmt.Sprintf("\033[33m Found %d files (Recursive). Use comma for AND logic. [Backsp] Clear\033[0m", len(state.FilteredFiles)))
	} else {
		sb.WriteString("\033[90m [/] Search (Example: 'main, .go, import')  [q] Back\033[0m")
	}

	fmt.Print(sb.String())
}

func RenderSettingsUI(width, height int, state *StateAplikasi, cfg KonfigurasiApp) {
	var sb strings.Builder
	sb.WriteString(cfg.Tema.Background)
	sb.WriteString("\033[H\033[2J") 
	cx := width / 2 - 25; cy := height / 2 - 8
	sb.WriteString(MoveCursorStr(cx, cy)); sb.WriteString("\033[33m=== KONFIGURASI ===")
	sb.WriteString(MoveCursorStr(cx, cy+2))
	p0 := "   "; if state.SettingsIdx == 0 { p0 = " ➤ " }
	sb.WriteString(fmt.Sprintf("%s%sDefault Editor: [%s] \033[30;47m[ ENTER ]\033[0m", cfg.Tema.Foreground, p0, cfg.EditorDefault))
	sb.WriteString(MoveCursorStr(cx, cy+3))
	p1 := "   "; if state.SettingsIdx == 1 { p1 = " ➤ " }
	isEdit1 := state.IsEditing && state.SettingsIdx == 1
	status1 := ""; if isEdit1 { status1 = " \033[31m[EDITING]\033[0m" }
	renderedText1 := RenderEditableText(cfg.ImagePath, state.EditCursorPos, isEdit1)
	sb.WriteString(fmt.Sprintf("%s%sImage Path: %s%s", cfg.Tema.Foreground, p1, renderedText1, status1))
	sb.WriteString(MoveCursorStr(cx, cy+4))
	p2 := "   "; if state.SettingsIdx == 2 { p2 = " ➤ " }
	isEdit2 := state.IsEditing && state.SettingsIdx == 2
	status2 := ""; if isEdit2 { status2 = " \033[31m[EDITING]\033[0m" }
	renderedText2 := RenderEditableText(cfg.FooterText, state.EditCursorPos, isEdit2)
	sb.WriteString(fmt.Sprintf("%s%sFooter Text: %s%s", cfg.Tema.Foreground, p2, renderedText2, status2))
	sb.WriteString(MoveCursorStr(cx, cy+6)); sb.WriteString("   --- Dialog Goldship ---")
	for i, dialog := range cfg.Dialogues {
		sb.WriteString(MoveCursorStr(cx, cy+7+i))
		currIdx := i + 3
		prefix := "   "; if state.SettingsIdx == currIdx { prefix = " ➤ " }
		isEditD := state.IsEditing && state.SettingsIdx == currIdx
		status := ""; if isEditD { status = " \033[31m[EDITING]\033[0m" }
		renderedDialog := RenderEditableText(dialog, state.EditCursorPos, isEditD)
		sb.WriteString(fmt.Sprintf("%s%sDialog %d: %s%s", cfg.Tema.Foreground, prefix, i+1, renderedDialog, status))
	}
	sb.WriteString(MoveCursorStr(cx, height-3))
	if state.IsEditing { sb.WriteString("\033[33m[←/→] Geser  [Ketik] Insert  [Backspace] Hapus  [ENTER] Simpan") } else { sb.WriteString("Navigasi: [j/k]. Ubah: [ENTER]. Kembali: [q]") }
	fmt.Print(sb.String())
}

func RenderAboutUI(width, height int, cfg KonfigurasiApp) {
	var sb strings.Builder
	sb.WriteString(cfg.Tema.Background)
	sb.WriteString("\033[H\033[2J")
	cx := width / 2 - 25; cy := height / 2 - 5
	sb.WriteString(MoveCursorStr(cx, cy)); sb.WriteString("\033[33m=== TENTANG GOLDFILE ===")
	sb.WriteString(MoveCursorStr(cx, cy+2)); sb.WriteString(fmt.Sprintf("%sGoldshin... Golshin...", cfg.Tema.Foreground))
	sb.WriteString(MoveCursorStr(cx, cy+3)); sb.WriteString("GoldFile adalah File-Manager seperti Superfile")
	sb.WriteString(MoveCursorStr(cx, cy+4)); sb.WriteString("Tapi lebih Simple, Newbie-Friendly dan ada Goldship-nya")
	sb.WriteString(MoveCursorStr(cx, cy+6)); sb.WriteString("\033[32mTekan tombol apa saja untuk kembali...")
	fmt.Print(sb.String())
}
