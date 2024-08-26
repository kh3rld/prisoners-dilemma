package common

import (
	"fmt"
	"math/rand"
	"net"
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
	SetName(name string)
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
