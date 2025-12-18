package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	config := InisialisasiConfig()
	state := InitState()

	err := EnableRawMode()
	if err != nil {
		fmt.Println("Error Raw Mode:", err)
		return
	}
	
	defer func() {
		DisableRawMode()
		
		fmt.Print("\033[2J\033[3J\033[H")
		
		fmt.Print("\033[?25h") // Show Cursor
		fmt.Println("GoldFile ditutup. Sampai jumpa Trainer-san! ✨")
	}()

	inputChan := make(chan []byte)
	go func() {
		buf := make([]byte, 3)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil { return }
			inputChan <- buf[:n]
		}
	}()

	ticker := time.NewTicker(200 * time.Millisecond) 
	running := true
	frame := 0

	for running {
		w, h := GetTerminalSize()
		
		RenderManager(w, h, frame, &state, config)

		select {
		case <-ticker.C:
			frame++
		case keyBytes := <-inputChan:
			strKey := string(keyBytes)
			if len(keyBytes) > 0 && keyBytes[0] == 3 {
				running = false
				break
			}

			switch state.Mode {
			case 0: running = state.HandleInputDashboard(strKey)
			case 1: state.HandleInputFileManager(strKey, &config)
			case 2: state.HandleInputSettings(strKey, &config)
			case 3: state.HandleInputAbout(strKey)
			}
		}
	}
}
