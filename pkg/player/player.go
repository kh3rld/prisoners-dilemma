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

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) SetTotalYears(years int) {
	p.TotalYears = years
}

func (p *Player) GetTotalYears() int {
	return p.TotalYears
}
