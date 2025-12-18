package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type KonfigurasiTema struct {
	Background    string
	Foreground    string
	SelectedBg    string
	SelectedFg    string
	Directory     string
	File          string
	Border        string
	StatusBarBg   string
}

type KonfigurasiApp struct {
	EditorDefault string  
	Tema          KonfigurasiTema
	
	ImagePath     string   
	FooterText    string   
	Dialogues     []string 
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "goldfile")

	
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

func InisialisasiConfig() KonfigurasiApp {
	if loadedCfg, err := LoadConfig(); err == nil {
		if loadedCfg.FooterText == "" {
			loadedCfg.FooterText = "[↑/↓] Navigasi  [ENTER] Pilih  [Tab] Preview"
		}
		return loadedCfg
	}

	tokyoNight := KonfigurasiTema{
		Background:    "\033[48;2;26;27;38m",
		Foreground:    "\033[38;2;192;202;245m",
		SelectedBg:    "\033[48;2;65;72;104m",
		SelectedFg:    "\033[38;2;122;162;247m",
		Directory:     "\033[38;2;122;162;247m",
		File:          "\033[38;2;154;165;206m",
		Border:        "\033[38;2;86;95;137m",
		StatusBarBg:   "\033[48;2;187;154;247m",
	}

	defaultCfg := KonfigurasiApp{
		EditorDefault: "nvim",
		Tema:          tokyoNight,
		ImagePath:     "goldship.png", 
		FooterText:    "[↑/↓] Navigasi  [ENTER] Pilih  [Tab] Preview",
		Dialogues: []string{
			"Nggak usah coding, Trainer-san...",
			"Mending scroll Fasnuk...",
			"Kopi mana kopi?",
		},
	}
	
	SaveConfig(&defaultCfg)
	return defaultCfg
}

func SaveConfig(cfg *KonfigurasiApp) error {
	path, err := GetConfigPath()
	if err != nil { return err }

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil { return err }
	
	return os.WriteFile(path, data, 0644)
}

func LoadConfig() (KonfigurasiApp, error) {
	var cfg KonfigurasiApp
	
	path, err := GetConfigPath()
	if err != nil { return cfg, err }

	data, err := os.ReadFile(path)
	if err != nil { return cfg, err }
	
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
