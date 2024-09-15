package ui

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/game"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
)

func GreenText(text string) string {
	return Green + text + Reset
}

func BlueText(text string) string {
	return Blue + text + Reset
}

func GetRounds() int {
	var rounds int
	fmt.Println("Enter the number of rounds: ")
	fmt.Scanln(&rounds)
	return rounds
}

func DisplayAvailableActions() {
	fmt.Println("Available Acions")
	for key, action := range common.Actions {
		fmt.Printf("%s: %s\n", key, action)
	}
}

func GetPlayerAction(name string) string {
	DisplayAvailableActions()
	fmt.Printf("Enter action for %s: ", name)
	var action string
	if _, err := fmt.Scanln(&action); err != nil {
		fmt.Println("Invalid input. Try again.")
		return GetPlayerAction(name)
	}
	if validAcion, err := common.ValidateAction(action); err == nil {
		return validAcion
	} else {
		fmt.Println("Invalid action. Try again.")
		return GetPlayerAction(name)
	}
}

func DisplayMessage(message string, colorFunc func(a ...interface{}) string) {
	fmt.Println(colorFunc(message))
}

func DisplayOutcome(g *game.Game) {
	p1, p2 := g.Players[0], g.Players[1]
	fmt.Println(CenterText(fmt.Sprintf("%s chose to %s.", p1.Name, p1.GetAction())))
	fmt.Println(CenterText(fmt.Sprintf("%s chose to %s.", p2.Name, p2.GetAction())))
	fmt.Println(CenterText(g.Result.Description))
	fmt.Println(CenterText(fmt.Sprintf("%s gets %d year%s in prison.", p1.Name, g.Result.Player1, pluralize(g.Result.Player1))))
	fmt.Println(CenterText(fmt.Sprintf("%s gets %d year%s in prison.", p2.Name, g.Result.Player2, pluralize(g.Result.Player2))))
	fmt.Println(CenterText("-----------------------------"))
}

func DisplayRoundSummary(round int, g *game.Game) {
	p1, p2 := g.Players[0], g.Players[1]
	fmt.Printf("Round %d Summary:\n", round)
	fmt.Printf("%s chose %s, %s chose %s\n", p1.Name, p1.GetAction(), p2.Name, p2.GetAction())
	fmt.Printf("Outcome: %s gets %d year%s, %s gets %d year%s\n",
		p1.Name, g.Result.Player1, pluralize(g.Result.Player1),
		p2.Name, g.Result.Player2, pluralize(g.Result.Player2))
}

func pluralize(years int) string {
	if years == 1 {
		return ""
	}
	return "s"
}
