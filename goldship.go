package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var AppLogoBig = []string{
	`   ______ ____  __   ____  _____ ____  __   ______`,
	`  / ____// __ \/ /  / __ \/ ____//  _/ / /  / ____/`,
	` / / __ / / / / /  / / / / /_    / /  / /  / __/   `,
	`/ /_/ // /_/ / /__/ /_/ / __/  _/ /  / /___/ /___  `,
	`\____/ \____/_____/____/_/    /___/ /_____/_____/  `,
}

var AppLogoMini = []string{
	`===  G  O  L  D  F  I  L  E  ===`,
}

func RenderDashboard(width, height int, frameCounter int, state *StateAplikasi, cfg KonfigurasiApp) {
	var sb strings.Builder
	
	sb.WriteString("\033[H\033[2J") 
	sb.WriteString(cfg.Tema.Background) 

	var selectedLogo []string
	
	if height < 35 || width < 65 {
		selectedLogo = AppLogoMini
	} else {
		selectedLogo = AppLogoBig
	}

	totalContentHeight := len(selectedLogo) + 15 + 2 + 4 + 1 // Logo + Img + Dialog + Menu + Footer
	startY := (height - totalContentHeight) / 2
	if startY < 1 { startY = 1 }

	for i, line := range selectedLogo {
		xPos := (width - len(line)) / 2
		if xPos < 1 { xPos = 1 }
		sb.WriteString(MoveCursorStr(xPos, startY+i))
		sb.WriteString(fmt.Sprintf("\033[1;33m%s\033[0m", line))
	}

	targetImgW := 30
	targetImgH := 15 

	if state.AsciiCache == "" || state.LastImageUsed != cfg.ImagePath {
		cmd := exec.Command("chafa", cfg.ImagePath, "--format=symbols", fmt.Sprintf("--size=%dx%d", targetImgW, targetImgH))
		out, err := cmd.Output()
		if err != nil {
			state.AsciiCache = fmt.Sprintf("\n   [Error]\n   %s", err.Error())
		} else {
			state.AsciiCache = string(out)
		}
		state.LastImageUsed = cfg.ImagePath
	}

	lines := strings.Split(state.AsciiCache, "\n")
	
	imgStartY := startY + len(selectedLogo) + 1

	for i, line := range lines {
		imgX := (width - targetImgW) / 2
		if imgX < 1 { imgX = 1 }

		sb.WriteString(MoveCursorStr(1, imgStartY+i))
		sb.WriteString(strings.Repeat(" ", width)) 
		sb.WriteString(MoveCursorStr(imgX, imgStartY+i))
		sb.WriteString(line)
	}

	dialogY := imgStartY + len(lines)
	sb.WriteString(MoveCursorStr(1, dialogY))
	sb.WriteString(strings.Repeat(" ", width))

	if len(cfg.Dialogues) > 0 {
		dialogIdx := (frameCounter / 10) % len(cfg.Dialogues)
		currentDialog := cfg.Dialogues[dialogIdx]
		dialX := (width - len(currentDialog)) / 2
		if dialX < 1 { dialX = 1 }
		sb.WriteString(MoveCursorStr(dialX, dialogY))
		sb.WriteString(fmt.Sprintf("\033[38;2;100;200;255m\"%s\"\033[0m", currentDialog))
	}

	menuStartY := dialogY + 2
	options := []string{"Buka File Manager", "Pengaturan (Settings)", "Tentang (About)", "Keluar"}
	
	remainingSpace := (height - 1) - menuStartY
	
	visibleMenuCount := remainingSpace
	if visibleMenuCount < 1 { visibleMenuCount = 1 }
	if visibleMenuCount > 4 { visibleMenuCount = 4 }

	startOpt := 0
	if state.DashboardIdx >= visibleMenuCount {
		startOpt = state.DashboardIdx - visibleMenuCount + 1
	}
	if startOpt < 0 { startOpt = 0 }
	if startOpt > len(options)-visibleMenuCount { startOpt = len(options)-visibleMenuCount }
	if startOpt < 0 { startOpt = 0 }

	for i := 0; i < visibleMenuCount; i++ {
		idx := startOpt + i
		if idx >= len(options) { break }

		opt := options[idx]
		menuY := menuStartY + i
		
		sb.WriteString(MoveCursorStr(1, menuY))
		sb.WriteString(strings.Repeat(" ", width))

		if i == 0 && startOpt > 0 {
			sb.WriteString(MoveCursorStr((width/2)-1, menuY-1))
			sb.WriteString("\033[90m▴\033[0m") 
		}
		if i == visibleMenuCount-1 && idx < len(options)-1 {
			sb.WriteString(MoveCursorStr((width/2)-1, menuY+1))
			sb.WriteString("\033[90m▾\033[0m") 
		}

		menuX := (width - len(opt)) / 2
		if menuX < 1 { menuX = 1 }
		sb.WriteString(MoveCursorStr(menuX-2, menuY))

		if idx == state.DashboardIdx {
			sb.WriteString(fmt.Sprintf("%s%s ➤ %s %s", cfg.Tema.SelectedBg, cfg.Tema.SelectedFg, opt, "\033[0m"))
		} else {
			sb.WriteString(fmt.Sprintf("%s   %s", cfg.Tema.Foreground, opt))
		}
	}

	footLen := len(cfg.FooterText)
	footX := (width - footLen) / 2
	if footX < 1 { footX = 1 }

	sb.WriteString(MoveCursorStr(footX, height-1))
	sb.WriteString(fmt.Sprintf("\033[33m%s", cfg.FooterText))

	fmt.Print(sb.String())
}
