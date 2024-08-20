package ai

import (
	"math/rand"
	"time"
)

type Opponent interface {
	ChooseAction(previousOpponentAction string) string
}

type TitForTat struct{}

func (t *TitForTat) ChooseAction(previousOpponentAction string) string {
	if previousOpponentAction == "" {
		return "1"
	}
	return previousOpponentAction
}

type Random struct{}

func (r *Random) ChooseAction(previousOpponentAction string) string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return "1"
	}
	return "2"
}
