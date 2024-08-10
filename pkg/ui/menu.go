package ui

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/term"
)

func DisplayHelp() {
	fmt.Println(CenterText("Prisoner's Dilemma Game Instructions:"))
	fmt.Println(CenterText("1. Each player can choose to cooperate or defect."))
	fmt.Println(CenterText("2. If both players cooperate, they each get 1 year."))
	fmt.Println(CenterText("3. If one defects and the other cooperates, the defector goes free and the cooperator gets 3 years."))
	fmt.Println(CenterText("4. If both defect, they each get 2 years."))
	fmt.Println(CenterText("Use '1' for cooperate, '2' for defect."))
}

func GetTerminalWidth() int {
	if runtime.GOOS == "windows" {
		return 80
	}

	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		return 80
	}
	return width
}

func CenterText(text string) string {
	width := GetTerminalWidth()
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%*s", padding+len(text), text)
}
