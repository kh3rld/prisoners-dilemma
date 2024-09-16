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

func RunGame(g *game.Game, rounds int, detailedSummaries bool) {
	for i := 1; i <= rounds; i++ {
		fmt.Printf("\nStarting Round %d\n", i)
		for _, p := range g.Players {
			p.SetAction(ui.GetPlayerAction(p.GetName()))
		}

		g.DetermineOutcome()

		ui.DisplayOutcome(g)

		if detailedSummaries {
			ui.DisplayRoundSummary(i, g)
		}
	}
}
