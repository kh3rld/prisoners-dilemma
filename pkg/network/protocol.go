package network

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type GameState struct {
	Round        int `json:"round"`
	Player1Score int `json:"player1_score"`
	Player2Score int `json:"player2_score"`
}

type RoundOutcome struct {
	Player1Points int `json:"player1_points"`
	Player2Points int `json:"player2_points"`
}

// Represents the current state of a player in the game
type PlayerState struct {
	Action string
	Score  int
}

// Represents the current game state including players and the round number
type Game struct {
	Player1 PlayerState
	Player2 PlayerState
	Round   int
	mu      sync.Mutex
}

// Global game state (could be encapsulated in a manager)
var game = Game{
	Player1: PlayerState{Action: "", Score: 0},
	Player2: PlayerState{Action: "", Score: 0},
	Round:   1,
}

func parseAction(message string) string {
	parts := strings.Split(message, ":")
	if len(parts) == 2 && parts[0] == "PLAYER_ACTION" {
		return parts[1]
	}
	fmt.Println("Invalid player action message format")
	return ""
}

// Parse the game state JSON from the message
func parseGameState(message string) (GameState, error) {
	var gs GameState
	if strings.HasPrefix(message, "GAME_STATE_UPDATE:") {
		jsonPart := strings.TrimPrefix(message, "GAME_STATE_UPDATE:")
		err := json.Unmarshal([]byte(jsonPart), &gs)
		if err != nil {
			fmt.Println("Error parsing game state:", err)
			return gs, err
		}
		return gs, nil
	}
	fmt.Println("Invalid game state message format")
	return gs, fmt.Errorf("invalid format")
}

// Parse the round outcome JSON from the message
func parseOutcome(message string) (RoundOutcome, error) {
	var outcome RoundOutcome
	if strings.HasPrefix(message, "ROUND_OUTCOME:") {
		jsonPart := strings.TrimPrefix(message, "ROUND_OUTCOME:")
		err := json.Unmarshal([]byte(jsonPart), &outcome)
		if err != nil {
			fmt.Println("Error parsing round outcome:", err)
			return outcome, err
		}
		return outcome, nil
	}

	fmt.Println("Invalid round outcome message format")
	return outcome, fmt.Errorf("invalid format")
}

// Process player action: Update the player's current round action
func processPlayerAction(playerID int, action string) {
	game.mu.Lock()
	defer game.mu.Unlock()

	if playerID == 1 {
		game.Player1.Action = action
		fmt.Printf("Player 1 action set to: %s\n", action)
	} else if playerID == 2 {
		game.Player2.Action = action
		fmt.Printf("Player 2 action set to: %s\n", action)
	}

	// Logic for resolving the round could go here if both players have acted
	if game.Player1.Action != "" && game.Player2.Action != "" {
		fmt.Println("Both players have acted. Proceed to resolving the round.")
		// TODO: Implement resolveRound()
	}
}

// Update the overall game state (round, scores)
func updateGameState(newState GameState) {
	game.mu.Lock()
	defer game.mu.Unlock()

	// Update the current round and scores
	game.Round = newState.Round
	game.Player1.Score = newState.Player1Score
	game.Player2.Score = newState.Player2Score

	fmt.Printf("Game state updated: Round %d, Player 1 score: %d, Player 2 score: %d\n", game.Round, game.Player1.Score, game.Player2.Score)
}

// Apply the round outcome to update scores based on the outcome
func applyRoundOutcome(outcome RoundOutcome) {
	game.mu.Lock()
	defer game.mu.Unlock()

	// Update player scores
	game.Player1.Score += outcome.Player1Points
	game.Player2.Score += outcome.Player2Points

	fmt.Printf("Round outcome applied. Player 1 points: %d, Player 2 points: %d\n", outcome.Player1Points, outcome.Player2Points)

	// Clear actions for the next round
	game.Player1.Action = ""
	game.Player2.Action = ""

	// Advance to the next round
	game.Round++
	fmt.Printf("Proceeding to round %d\n", game.Round)
}
