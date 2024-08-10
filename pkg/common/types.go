package common

type GameInterface interface {
	GetOutcome() Outcome
}

type Outcome struct {
	Player1     int
	Player2     int
	Description string
}
