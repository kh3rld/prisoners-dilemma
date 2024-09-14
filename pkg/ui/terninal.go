package ui

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/buger/goterm"
	"golang.org/x/term"
)

const DefaultWidth = 80

func GetTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err == nil {
		return width
	}
	log.Printf("Failed to get terminal size with term.GetSize(): %v", err)

	// If the above fails, attempt fallback methods.
	if runtime.GOOS == "windows" {
		return getWindowsFallbackWidth()
	}

	// Check for $COLUMNS environment variable
	if columnsEnv := os.Getenv("COLUMNS"); columnsEnv != "" {
		if width, err := strconv.Atoi(columnsEnv); err == nil && width > 0 {
			return width
		}
	}

	if fallbackWidth := goterm.Width(); fallbackWidth > 0 {
		return fallbackWidth
	}

	return promptForTerminalWidth()
}

func getWindowsFallbackWidth() int {
	type winsize struct {
		ws_row    uint16
		ws_col    uint16
		ws_xpixel uint16
		ws_ypixel uint16
	}

	fd := int(os.Stdout.Fd())

	const TIOCGWINSZ = 0x5413

	var ws winsize

	// Call ioctl to get the terminal size
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), TIOCGWINSZ, uintptr(unsafe.Pointer(&ws)))
	if errno != 0 {
		log.Printf("Failed to get terminal size: %v", errno)
		return DefaultWidth
	}

	width := int(ws.ws_col)
	log.Printf("Detected terminal width: %d", width)
	return width
}

// Prompt the user for the terminal width if automatic methods fail
func promptForTerminalWidth() int {
	fmt.Printf("Unable to detect terminal width. Please enter terminal width (default: %d): ", DefaultWidth)

	var input string
	fmt.Scanln(&input)

	width, err := strconv.Atoi(input)
	if err != nil || width <= 0 {
		log.Printf("Invalid input. Falling back to default width: %d", DefaultWidth)
		return DefaultWidth
	}
	return width
}
