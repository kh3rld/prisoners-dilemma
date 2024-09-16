package ui

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"golang.org/x/term"
)

func DisplayArt() {
	art, err := LoadArt()
	if err != nil {
		fmt.Printf("erro loading %v", err)
	}
	fmt.Println(CenterText(BlueText(art)))
	fmt.Println(CenterText(GreenText("Developed by Kherld Hussein")))
	fmt.Println(CenterText(GreenText("Special thanks to the Go community")))
	fmt.Println(CenterText(GreenText("and open source contributors.")))
	fmt.Println()
}

func LoadArt() (string, error) {
	file, err := os.Open("../configs/art.txt")
	if err != nil {
		return "", fmt.Errorf("error opening art: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data strings.Builder

	for scanner.Scan() {
		data.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading art: %v", err)
	}

	return data.String(), nil
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

func CenterText(text string) string {
	width := GetTerminalWidth()
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%*s", padding+len(text), text)
}
