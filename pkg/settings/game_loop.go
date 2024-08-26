package settings

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/network"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
)

var player1Name, player2Name string

func SetPlayers() (common.PlayerInterface, common.PlayerInterface) {
	p1 := &player.Player{Name: player1Name}
	p2 := &player.Player{Name: player2Name}

	if player1Name == "" {
		fmt.Print("Enter the name of Player 1: ")
		fmt.Scanln(&player1Name)
		p1.SetName(player1Name)
	}

	if player2Name == "" {
		fmt.Print("Enter the name of Player 2: ")
		fmt.Scanln(&player2Name)
		p2.SetName(player2Name)
	}

	return p1, p2
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

func GameLoop(conn interface{}, player1, player2 *player.Player, rounds int, detailedSummaries bool) {
	game := &game.Game{
		Player1: player1,
		Player2: player2,
		Rounds:  rounds,
	}

	if server, ok := conn.(*network.Server); ok {
		game.Server = server
		game.StartGameServer(detailedSummaries)
	} else if client, ok := conn.(*network.Client); ok {
		game.StartGameClient(client, detailedSummaries)
	} else {
		fmt.Println("Starting a local game...")
		for i := 1; i <= rounds; i++ {
			fmt.Printf("\nStarting Round %d\n", i)
			player1.SetAction(ui.GetPlayerAction(player1.GetName()))
			player2.SetAction(ui.GetPlayerAction(player2.GetName()))
			game.DetermineOutcome()
			ui.DisplayOutcome(*player1, *player2, *game)
			ui.DisplayRoundSummary(i, *player1, *player2, *game)
		}
		return
	}
}
