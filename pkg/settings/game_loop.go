package settings

import (
	"errors"
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
)

func SetPlayers() (*player.Player, *player.Player, error) {
	p1 := &player.Player{}
	p2 := &player.Player{}

	fmt.Print("Enter the name of Player 1: ")
	if _, err := fmt.Scanln(&p1.Name); err != nil {
		return nil, nil, errors.New("failed to read player 1 name")
	}

	fmt.Print("Enter the name of Player 2: ")
	if _, err := fmt.Scanln(&p2.Name); err != nil {
		return nil, nil, errors.New("failed to read player 2 name")
	}

	return p1, p2, nil
}

func RunGame(p1, p2 *player.Player, rounds int, detailedSummaries bool) {
	game := &game.Game{Player1: p1, Player2: p2, Rounds: rounds}
	for i := 1; i <= rounds; i++ {
		fmt.Printf("\nStarting Round %d\n", i)
		p1.SetAction(ui.GetPlayerAction(p1.Name))
		p2.SetAction(ui.GetPlayerAction(p2.Name))
		game.DetermineOutcome()
		ui.DisplayOutcome(*p1, *p2, *game)
		if detailedSummaries {
			ui.DisplayRoundSummary(i, *p1, *p2, *game)
		}
	}
}
