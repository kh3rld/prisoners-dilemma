package game

import (
	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
)

type Game struct {
	Player1 player.Player
	Player2 player.Player
	Rounds  int
	Result  common.Outcome
}

func (g *Game) DetermineOutcome() {
	action1 := g.Player1.GetAction()
	action2 := g.Player2.GetAction()

	switch {
	case action1 == "cooperate" && action2 == "cooperate":
		g.Result = common.Outcome{Player1: 1, Player2: 1, Description: "Both players cooperated. Each gets 1 year in prison."}
	case action1 == "defect" && action2 == "cooperate":
		g.Result = common.Outcome{Player1: 0, Player2: 3, Description: "Player 1 defects. Player 1 goes free, Player 2 gets 3 years in prison."}
	case action1 == "cooperate" && action2 == "defect":
		g.Result = common.Outcome{Player1: 3, Player2: 0, Description: "Player 2 defects. Player 2 goes free, Player 1 gets 3 years in prison."}
	case action1 == "defect" && action2 == "defect":
		g.Result = common.Outcome{Player1: 2, Player2: 2, Description: "Both players defect. Each gets 2 years in prison."}
	}

	g.Player1.TotalYears += g.Result.Player1
	g.Player2.TotalYears += g.Result.Player2
}
