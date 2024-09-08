package main

import (
	"fmt"
	"os"

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
	player1, player2 := settings.SetPlayers()
	rounds, detailedSummaries := settings.GetUserSettings()

	p1, ok1 := player1.(*player.Player)
	p2, ok2 := player2.(*player.Player)

	if !ok1 || !ok2 {
		fmt.Println("Error: Players must be of type *player.Player")
		return
	}

	p1.SetName(player1.GetName())
	p2.SetName(player2.GetName())

	settings.GameLoop(nil, p1, p2, rounds, detailedSummaries)
}
