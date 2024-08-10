package main

import (
	"fmt"
	"os"

	"github.com/kh3rld/prisoners-dilemma/pkg/settings"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
	"github.com/kh3rld/prisoners-dilemma/pkg/utils"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func DisplayArt() {
	fmt.Println(ui.CenterText(ui.GreenText("===================================")))
	fmt.Println(ui.CenterText(ui.BlueText("  PRISONER'S DILEMMA GAME")))
	fmt.Println(ui.CenterText(ui.GreenText("===================================")))
	fmt.Println()
	fmt.Println(ui.CenterText(ui.GreenText("Developed by Kherld Hussein")))
	fmt.Println(ui.CenterText(ui.GreenText("Special thanks to the Go community")))
	fmt.Println(ui.CenterText(ui.GreenText("and open source contributors.")))
	fmt.Println()
}

func DisplayMenu() {
	ClearScreen()
	DisplayArt()

	for {
		fmt.Println(ui.CenterText("1. Start Game"))
		fmt.Println(ui.CenterText("2. View Instructions"))
		fmt.Println(ui.CenterText("3. Quit"))

		choice := utils.GetValidatedChoice("Enter your choice (1/2/3): ", []string{"1", "2", "3"})

		switch choice {
		case "1":
			utils.ShowProgress("Starting game...")
			settings.GameLoop()
		case "2":
			ui.DisplayHelp()
		case "3":
			if utils.ConfirmExit() {
				ShowCursor()
				fmt.Println(ui.CenterText("Thank you for playing!"))
				os.Exit(0)
			}
		}
	}

}
