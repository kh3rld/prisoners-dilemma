package ui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
)

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRounds() int {
	var rounds int
	fmt.Println("Enter the number of rounds: ")
	fmt.Scanln(&rounds)
	return rounds
}

func GetPlayerAction(name string) string {
	strategies := []string{"cooperate", "defect"}

	var action string

	fmt.Printf("%s, Choose your action (1: cooperate / 2: defect): ", name)
	fmt.Scanln(&action)

	action = strings.ToLower(action)

	if action != "cooperate" && action != "defect" {
		action = strategies[src.Intn(len(strategies))]
		fmt.Printf("Invalid Input, randomly assigning action %s\n", action)
	}

	return action
}

func DisplayOutcome(player1 player.Player, player2 player.Player, outcome game.Game) {
	fmt.Printf("%s chose to %s.\n", player1.Name, player1.Action)
	fmt.Printf("%s chose to %s.\n", player2.Name, player2.Action)
	fmt.Println(outcome.Result.Description)
	fmt.Printf("%s gets %d years in prison.\n", player1.Name, outcome.Result.Player1)
	fmt.Printf("%s gets %d years in prison.\n", player2.Name, outcome.Result.Player2)
	fmt.Println("-----------------------------")
}

func DisplayFinalResults(player1 player.Player, player2 player.Player) {
	fmt.Println("Game Over!")
	fmt.Printf("Total years in prison for %s: %d\n", player1.Name, player1.TotalYears)
	fmt.Printf("Total years in prison for %s: %d\n", player2.Name, player2.TotalYears)

	if player1.TotalYears < player2.TotalYears {
		fmt.Printf("%s wins by cooperating more!\n", player1.Name)
	} else if player1.TotalYears > player2.TotalYears {
		fmt.Printf("%s wins by cooperating more!\n", player2.Name)
	} else {
		fmt.Println("It's a tie! Both players served the same amount of time.")
	}
}
