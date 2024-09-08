package ui

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
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

func GetPlayerAction(name string) string {
	var choice string
	fmt.Printf("%s, Choose your action (1: cooperate / 2: defect): ", name)
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		return "cooperate"
	case "2":
		return "defect"
	default:
		action := common.GetRandomAction()
		fmt.Printf("Invalid input, randomly assigning action: %s\n", action)
		return action
	}
}

func DisplayFinalResults(p1 player.Player, p2 player.Player) {
	fmt.Println("Game Over!")
	fmt.Printf("Total years in prison for %s: %d\n", p1.Name, p1.TotalYears)
	fmt.Printf("Total years in prison for %s: %d\n", p2.Name, p2.TotalYears)

	if p1.TotalYears < p2.TotalYears {
		fmt.Printf("%s wins by cooperating more!\n", p1.Name)
	} else if p1.TotalYears > p2.TotalYears {
		fmt.Printf("%s wins by cooperating more!\n", p2.Name)
	} else {
		fmt.Println("It's a tie! Both players served the same amount of time.")
	}
}

func DisplayMessage(message string, colorFunc func(a ...interface{}) string) {
	fmt.Println(colorFunc(message))
}

func DisplayGameStartMessage() {
	fmt.Println(GreenText("Welcome to the Prisoner's Dilemma Game!"))
	fmt.Println(BlueText("Developed by Kherld Hussein"))
	fmt.Println("Special thanks to the Go community and open source contributors.")
}

func DisplayOutcome(p1 player.Player, p2 player.Player, game game.Game) {
	fmt.Println(CenterText(fmt.Sprintf("%s chose to %s.", p1.Name, p1.Action)))
	fmt.Println(CenterText(fmt.Sprintf("%s chose to %s.", p2.Name, p2.Action)))
	fmt.Println(CenterText(game.Result.Description))
	fmt.Println(CenterText(fmt.Sprintf("%s gets %d years in prison.", p1.Name, game.Result.Player1)))
	fmt.Println(CenterText(fmt.Sprintf("%s gets %d years in prison.", p2.Name, game.Result.Player2)))
	fmt.Println(CenterText("-----------------------------"))
}

func DisplayRoundSummary(round int, p1, p2 player.Player, outcome game.Game) {
	fmt.Printf("Round %d Summary:\n", round)
	fmt.Printf("%s chose %s, %s chose %s\n", p1.Name, p1.Action, p2.Name, p2.Action)
	fmt.Printf("Outcome: %s gets %d year%s, %s gets %d year%s\n", p1.Name, outcome.Result.Player1, pluralize(outcome.Result.Player1), p2.Name, outcome.Result.Player2, pluralize(outcome.Result.Player2))
}

func pluralize(years int) string {
	if years == 1 {
		return ""
	}
	return "s"
}
