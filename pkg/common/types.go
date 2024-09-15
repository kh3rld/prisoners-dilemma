package common

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

const (
	ActionCooperate = "cooperate"
	ActionDefect    = "defect"
)

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

var Actions = map[string]string{
	"C": "cooperate",
	"D": "defect",
}

type PlayerInterface interface {
	SetAction(action string)
	GetAction() string
	GetName() string
	SetTotalYears(years int)
	GetTotalYears() int
	SetName(name string)
}

type Outcome struct {
	Player1     int
	Player2     int
	Description string
}

func ValidateAction(action string) (string, error) {
	action = strings.ToUpper(action)
	if validAction, exists := Actions[action]; exists {
		return validAction, nil
	}
	return "", fmt.Errorf("invalid action")
}

func AddOrModifyAction(key, action string) {
	Actions[key] = action
}

func GetRandomAction() string {
	actions := []string{"cooperate", "defect"}
	return actions[src.Intn(len(actions))]
}

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no IP found")
}
