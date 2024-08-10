package settings

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
	"github.com/kh3rld/prisoners-dilemma/pkg/utils"
)

var player1Name, player2Name string

func SetPlayers() (player.Player, player.Player) {
	var player1, player2 player.Player

	if player1Name == "" {
		fmt.Print("Enter the name of Player 1: ")
		fmt.Scanln(&player1Name)
	}
	player1 = player.Player{Name: player1Name}

	if player2Name == "" {
		fmt.Print("Enter the name of Player 2: ")
		fmt.Scanln(&player2Name)
	}
	player2 = player.Player{Name: player2Name}

	return player1, player2
}

func PlayMultipleRounds(rounds int, player1, player2 *player.Player) error {
	g := game.Game{Player1: player1, Player2: player2, Rounds: rounds}

	for i := 1; i <= rounds; i++ {
		fmt.Printf("\nStarting Round %d\n", i)

		player1.SetAction(ui.GetPlayerAction(player1.GetName()))
		player2.SetAction(ui.GetPlayerAction(player2.GetName()))

		g.DetermineOutcome()

		ui.DisplayOutcome(*player1, *player2, g)

		ui.DisplayRoundSummary(i, *player1, *player2, g)
	}
	return nil
}
func GameLoop() {
	for {
		player1, player2 := SetPlayers()

		rounds, showSummary := GetUserSettings()

		if err := PlayMultipleRounds(rounds, &player1, &player2); err != nil {
			fmt.Println("Error during game:", err)
			continue
		}

		if showSummary {
			ui.DisplayFinalResults(player1, player2)
		}

		if !utils.ConfirmExit() {
			break
		}
	}
}
