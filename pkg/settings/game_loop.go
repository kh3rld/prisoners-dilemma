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
	p1 := &player.Player{}
	p2 := &player.Player{}

	if player1Name == "" {
		fmt.Print("Enter the name of Player 1: ")
		fmt.Scanln(&player1Name)
	}
	// player1 = player.Player{Name: player1Name}

	if player2Name == "" {
		fmt.Print("Enter the name of Player 2: ")
		fmt.Scanln(&player2Name)
	}
	// player2 = player.Player{Name: player2Name}

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

func GameLoop(conn interface{}, player1, player2 common.PlayerInterface) {
	for {
		if conn == nil {
			player1, player2 = SetPlayers()
		}

		rounds, showSummary := GetUserSettings()

		for round := 1; round <= rounds; round++ {
			var action1, action2 string
			if conn == nil {
				action1 = player1.GetAction()
				action2 = player2.GetAction()
			} else {
				switch c := conn.(type) {
				case *network.Server:
					action1 = c.ReceiveAction(c.Clients[0])
					action2 = c.ReceiveAction(c.Clients[1])
				case *network.Client:
					c.SendAction(player1.GetAction())
					action1 = c.ReceiveResult()
					c.SendAction(player2.GetAction())
					action2 = c.ReceiveResult()
				}
			}

			player1.SetAction(action1)
			player2.SetAction(action2)

			g := game.Game{Player1: player1, Player2: player2}
			g.DetermineOutcome()
			outcome := g.Result.Description

			if conn == nil {
				ui.DisplayRoundSummary(round, *player1.(*player.Player), *player2.(*player.Player), g)
			} else {
				switch c := conn.(type) {
				case *network.Server:
					c.SendResult(c.Clients[0], outcome)
					c.SendResult(c.Clients[1], outcome)
				case *network.Client:
					c.ReceiveResult()
				}
			}
		}
		if !showSummary {
			break
		}
	}
}
