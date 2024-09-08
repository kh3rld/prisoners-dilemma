package game

import (
	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
)

type Game struct {
	Player1 *player.Player
	Player2 *player.Player
	Rounds  int
	Result  common.Outcome
}

func (g *Game) DetermineOutcome() {
	a1 := g.Player1.GetAction()
	a2 := g.Player2.GetAction()

	switch {
	case a1 == common.ActionCooperate && a2 == common.ActionCooperate:
		g.Result = common.Outcome{Player1: 1, Player2: 1, Description: "Both players cooperated. Each gets 1 year in prison."}
	case a1 == common.ActionDefect && a2 == common.ActionCooperate:
		g.Result = common.Outcome{Player1: 0, Player2: 3, Description: "Player 1 defects. Player 1 goes free, Player 2 gets 3 years in prison."}
	case a1 == common.ActionCooperate && a2 == common.ActionDefect:
		g.Result = common.Outcome{Player1: 3, Player2: 0, Description: "Player 2 defects. Player 2 goes free, Player 1 gets 3 years in prison."}
	case a1 == common.ActionDefect && a2 == common.ActionDefect:
		g.Result = common.Outcome{Player1: 2, Player2: 2, Description: "Both players defect. Each gets 2 years in prison."}
	}

	g.Player1.SetTotalYears(g.Player1.GetTotalYears() + g.Result.Player1)
	g.Player2.SetTotalYears(g.Player2.GetTotalYears() + g.Result.Player2)

}

func (g *Game) PlayRound(round int, player1, player2 *player.Player) common.Outcome {
	player1Action := player1.GetAction()
	if player1Action == "" {
		player1Action = common.GetRandomAction()
		player1.SetAction(player1Action)
	}

	player2Action := player2.GetAction()
	if player2Action == "" {
		player2Action = common.GetRandomAction()
		player2.SetAction(player2Action)
	}

	g.DetermineOutcome()

	return g.Result
}
