package common

import (
	"math/rand"
	"strings"
	"time"
)

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

type PlayerInterface interface {
	SetAction(action string)
	GetAction() string
	GetName() string
	SetTotalYears(years int)
	GetTotalYears() int
}

type Outcome struct {
	Player1     int
	Player2     int
	Description string
}

func ValidateAction(action string) string {
	action = strings.ToLower(action)
	if action == "cooperate" || action == "defect" {
		return action
	}
	return "cooperate"
}

func GetRandomAction() string {
	actions := []string{"cooperate", "defect"}
	return actions[src.Intn(len(actions))]
}
