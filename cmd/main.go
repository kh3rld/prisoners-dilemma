package main

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
)

func main() {
	player1 := player.Player{Name: "Kherld"}
	player2 := player.Player{Name: "Ilara"}

	game := game.Game{Player1: player1, Player2: player2}

	game.Player1.SetAction("cooperate")
	game.Player2.SetAction("defect")

	game.DetermineOutcome()

	outcome := game.Result
	fmt.Println(outcome.Description)
	fmt.Printf("%s gets %d years in prison.\n", game.Player1.Name, outcome.Player1)
	fmt.Printf("%s gets %d years in prison.\n", game.Player2.Name, outcome.Player2)
}
