package player

import "github.com/kh3rld/prisoners-dilemma/pkg/ai"

type AIPlayer struct {
	name       string
	ai         ai.Opponent
	action     string
	totalYears int
}

func NewAIPlayer(name string, ai ai.Opponent) *AIPlayer {
	return &AIPlayer{name: name, ai: ai}
}

func (p *AIPlayer) GetAction() string {
	if p.action == "" {
		p.action = p.ai.ChooseAction("")
	}
	return p.action
}

func (p *AIPlayer) SetAction(action string) {
	p.action = action
}

func (p *AIPlayer) GetName() string {
	return p.name
}

func (p *AIPlayer) GetTotalYears() int {
	return p.totalYears
}

func (p *AIPlayer) SetTotalYears(years int) {
	p.totalYears = years
}
