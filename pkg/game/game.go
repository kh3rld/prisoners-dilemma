package game

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
)

type Game struct {
	Players []*player.Player
	Rounds  int
	Config  *GameConfig
	Result  common.Outcome
}

type Rule struct {
	Action1     string `json:"action1"`
	Action2     string `json:"action2"`
	Outcome1    int    `json:"outcome1"`
	Outcome2    int    `json:"outcome2"`
	Description string `json:"description"`
}

type GameConfig struct {
	Rules []Rule `json:"rules"`
}

func LoadGameConfig(file string) (*GameConfig, error) {
	projectRoot, err := FindProjectRoot()
	if err != nil {
		fmt.Println("Error finding project root:", err)
		return nil, err
	}

	configFilePath := filepath.Join(projectRoot, file)
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config GameConfig
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (g *Game) DetermineOutcome() {
	if len(g.Players) < 2 {
		g.Result.Description = "Not enough players to determine outcome."
		return
	}

	// Get actions of both players
	p1 := g.Players[0]
	p2 := g.Players[1]

	action1 := p1.GetAction()
	action2 := p2.GetAction()

	// Look for a matching rule
	for _, rule := range g.Config.Rules {
		if (rule.Action1 == action1 && rule.Action2 == action2) ||
			(rule.Action1 == action2 && rule.Action2 == action1) {

			// Determine which outcome to apply to each player
			if rule.Action1 == action1 && rule.Action2 == action2 {
				p1.SetTotalYears(p1.GetTotalYears() + rule.Outcome1)
				p2.SetTotalYears(p2.GetTotalYears() + rule.Outcome2)
			} else {
				p1.SetTotalYears(p1.GetTotalYears() + rule.Outcome2)
				p2.SetTotalYears(p2.GetTotalYears() + rule.Outcome1)
			}

			// Set the result description
			g.Result.Description = rule.Description

			// Update the result struct with the players' total prison years
			g.Result.Player1 = p1.GetTotalYears()
			g.Result.Player2 = p2.GetTotalYears()

			return
		}
	}

	// No matching rule found
	g.Result.Description = "No matching rules for these actions."
}

func FindProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse up the directory tree until we find go.mod
	for {
		modPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(modPath); !os.IsNotExist(err) {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}

func (g *Game) PlayRound(round int, player1, player2 *player.Player) common.Outcome {
	player1Action := player1.GetAction()
	if player1Action == "" {
		player1Action = common.GetRandomAction()
		player1.SetAction(player1Action)
	}

	player2Action := player2.GetAction()
	if player2Action == "" {
		player2Action = common.GetRandomAction()
		player2.SetAction(player2Action)
	}

	g.DetermineOutcome()

	return g.Result
}
