package ui

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/term"
)

func DisplayArt() {
	fmt.Println(CenterText(GreenText("===================================")))
	fmt.Println(CenterText(BlueText("  PRISONER'S DILEMMA GAME")))
	fmt.Println(CenterText(GreenText("===================================")))
	fmt.Println()
	fmt.Println(CenterText(GreenText("Developed by Kherld Hussein")))
	fmt.Println(CenterText(GreenText("Special thanks to the Go community")))
	fmt.Println(CenterText(GreenText("and open source contributors.")))
	fmt.Println()
}

func DisplayHelp() {

	fmt.Println(CenterText("Instructions:"))
	fmt.Println(CenterText("1. Choose whether to play locally or over the network."))
	fmt.Println(CenterText("2. If playing over the network, you can host or join a game."))
	fmt.Println(CenterText("3. After establishing the connection, you'll see your opponent's name."))
	fmt.Println(CenterText("4. Follow the prompts to make your choices."))
	fmt.Println(CenterText("5. The game will display the outcome after each round."))

	fmt.Println()
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
