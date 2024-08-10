package main

import (
	"fmt"

	"github.com/kh3rld/prisoners-dilemma/pkg/game"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
)

func setPlayers() (player.Player, player.Player) {
	var player1Name, player2Name string

	fmt.Print("Enter the name of Player 1: ")
	fmt.Scanln(&player1Name)

	fmt.Print("Enter the name of Player 2: ")
	fmt.Scanln(&player2Name)

	player1 := player.Player{Name: player1Name}
	player2 := player.Player{Name: player2Name}

	return player1, player2
}

func playRound(round int, player1, player2 *player.Player) {
	fmt.Printf("Round %d\n", round)

	player1.SetAction(ui.GetPlayerAction(player1.Name))
	player2.SetAction(ui.GetPlayerAction(player2.Name))

	g := game.Game{Player1: *player1, Player2: *player2}

	g.DetermineOutcome()

	ui.DisplayOutcome(g.Player1, g.Player2, g)
}

func main() {
	player1, player2 := setPlayers()

	rounds := ui.GetRounds()

	for i := 1; i < rounds; i++ {
		playRound(i, &player1, &player2)
	}

	ui.DisplayFinalResults(player1, player2)
}
