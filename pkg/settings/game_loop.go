package settings

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
)

func SetPlayers() (common.PlayerInterface, common.PlayerInterface) {
	p1 := &player.Player{}
	p2 := &player.Player{}

	fmt.Print("Enter the name of Player 1: ")
	fmt.Scanln(&p1.Name)

	fmt.Print("Enter the name of Player 2: ")
	fmt.Scanln(&p2.Name)

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

	fmt.Println("Starting a local game...")
	for i := 1; i <= rounds; i++ {
		fmt.Printf("\nStarting Round %d\n", i)
		player1.SetAction(ui.GetPlayerAction(player1.GetName()))
		player2.SetAction(ui.GetPlayerAction(player2.GetName()))
		game.DetermineOutcome()
		ui.DisplayOutcome(*player1, *player2, *game)
		ui.DisplayRoundSummary(i, *player1, *player2, *game)
	}

}
