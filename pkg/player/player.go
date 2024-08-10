package player

type Player struct {
	Name       string
	Action     string
	TotalYears int
}

func (p *Player) SetAction(action string) {
	p.Action = action
}

func (p *Player) GetAction() string {
	return p.Action
}
