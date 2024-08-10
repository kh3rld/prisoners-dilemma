package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
)

func GetValidatedChoice(prompt string, validChoices []string) string {
	var choice string
	for {
		fmt.Print(ui.CenterText(prompt))
		fmt.Scanln(&choice)
		choice = strings.TrimSpace(choice)
		for _, v := range validChoices {
			if choice == v {
				return choice
			}
		}
		fmt.Println(ui.CenterText("Invalid choice. Please try again."))
	}
}

func ConfirmExit() bool {
	var response string
	fmt.Print(ui.CenterText("Are you sure you want to quit? (y/n): "))
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}

func ShowProgress(message string) {
	fmt.Print(ui.CenterText(message))
	spinner := []string{"|", "/", "-", "\\"}
	for i := 0; i < 20; i++ {
		fmt.Print("\r", ui.CenterText(message), spinner[i%len(spinner)])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Print("\r", ui.CenterText(message), " Done!\n")
}
