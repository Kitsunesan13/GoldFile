package main

import (
	"fmt"
	"strings"
	"syscall"
	"unicode"
	"unsafe"
)

var origTermState *syscall.Termios

type UkuranTerminal struct {
	Baris uint16
	Kolom uint16
	X     uint16
	Y     uint16
}

func GetTerminalSize() (int, int) {
	var sz UkuranTerminal
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.Kolom), int(sz.Baris)
}

func EnableRawMode() error {
	var oldState syscall.Termios
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&oldState))); err != 0 {
		return err
	}
	if origTermState == nil { origTermState = &oldState }
	
	newState := oldState
	newState.Iflag &^= syscall.IXON | syscall.ICRNL
	newState.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.ISIG | syscall.IEXTEN
	
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&newState)))
	
	fmt.Print("\033[?1049h\033[?1000h\033[H\033[2J") 
	return nil
}

func DisableRawMode() {
	if origTermState != nil {

		fmt.Print("\033[?1000l\033[?1049l") 
		
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(origTermState)))
	}
}

func ClearScreen() { fmt.Print("\033[H\033[2J") }
func MoveCursorStr(x, y int) string { return fmt.Sprintf("\033[%d;%dH", y, x) }
func MoveCursor(x, y int) { fmt.Printf("\033[%d;%dH", y, x) }

var keywords = map[string]bool{
	"func": true, "package": true, "import": true, "return": true,
	"if": true, "else": true, "for": true, "range": true, "switch": true, "case": true,
	"var": true, "const": true, "type": true, "struct": true, "interface": true,
	"true": true, "false": true, "nil": true, "int": true, "string": true, "bool": true,
	"class": true, "public": true, "private": true, "void": true, "def": true,
}

func HighlightCode(code string) string {
	var sb strings.Builder
	runes := []rune(code)
	length := len(runes)
	i := 0
	for i < length {
		char := runes[i]
		if i+1 < length && char == '/' && runes[i+1] == '/' {
			sb.WriteString("\033[90m"); for i < length { sb.WriteRune(runes[i]); i++ }; sb.WriteString("\033[0m"); continue
		}
		if char == '"' || char == '`' || char == '\'' {
			quote := char; sb.WriteString("\033[32m"); sb.WriteRune(char); i++
			for i < length { curr := runes[i]; sb.WriteRune(curr); if curr == quote && runes[i-1] != '\\' { i++; break }; i++ }
			sb.WriteString("\033[0m"); continue
		}
		if unicode.IsLetter(char) || char == '_' {
			start := i; for i < length && (unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]) || runes[i] == '_') { i++ }
			word := string(runes[start:i])
			if keywords[word] { sb.WriteString("\033[35m" + word + "\033[0m") } else { sb.WriteString("\033[37m" + word + "\033[0m") }
			continue
		}
		if unicode.IsDigit(char) {
			sb.WriteString("\033[33m"); for i < length && unicode.IsDigit(runes[i]) { sb.WriteRune(runes[i]); i++ }; sb.WriteString("\033[0m"); continue
		}
		if strings.ContainsRune(":=+-*/%&|<>!(){}[].,", char) { sb.WriteString("\033[36m"); sb.WriteRune(char); sb.WriteString("\033[0m"); i++; continue }
		sb.WriteRune(char); i++
	}
	return sb.String()
}

func TruncateAnsi(str string, maxLen int) string {
	if maxLen <= 0 { return "" }
	var sb strings.Builder
	visibleLen := 0
	inAnsi := false
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == '\033' { inAnsi = true; sb.WriteRune(r); continue }
		if inAnsi { sb.WriteRune(r); if r == 'm' { inAnsi = false }; continue }
		if visibleLen >= maxLen { break }
		if r == '\t' { sb.WriteString("  "); visibleLen += 2 } else { sb.WriteRune(r); visibleLen++ }
	}
	sb.WriteString("\033[0m")
	return sb.String()
}
