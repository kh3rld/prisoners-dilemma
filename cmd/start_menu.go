package main

import (
	"fmt"
	"os"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
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

func DisplayMenu() {
	ClearScreen()
	ui.DisplayArt()
	handleMenuChoice()
}

func handleMenuChoice() {
	for {
		fmt.Println(ui.CenterText("1. Play Locally"))
		fmt.Println(ui.CenterText("2. Play Over Network"))
		fmt.Println(ui.CenterText("3. View Instructions"))
		fmt.Println(ui.CenterText("4. Quit"))

		choice := utils.GetValidatedChoice("Enter your choice (1/2/3/4): ", []string{"1", "2", "3", "4"})

		switch choice {
		case "1":
			StartLocalGame()
		case "2":
		// TODO: Implement 	StartNetworkGame()
		case "3":
			ui.DisplayHelp()
		case "4":
			if utils.ConfirmExit() {
				ShowCursor()
				fmt.Println(ui.CenterText("Thank you for playing!"))
				os.Exit(0)
			}
		}
	}

}

func StartLocalGame() {
	utils.ShowProgress("Starting game...")
	configFilePath := "game_rules.json"

	// Load game config
	config, err := game.LoadGameConfig(configFilePath)
	if err != nil {
		fmt.Println("Error loading game config:", err)
		return
	}

	p1, p2, err := settings.SetPlayers()
	if err != nil {
		fmt.Println("Error: Could not set up players.")
		return
	}

	rounds, detailedSummaries := settings.GetUserSettings()

	p1.SetName(p1.GetName())
	p2.SetName(p2.GetName())

	g := game.Game{
		Players: []*player.Player{p1, p2},
		Rounds:  rounds,
		Config:  config,
	}

	settings.RunGame(&g, rounds, detailedSummaries)
}
