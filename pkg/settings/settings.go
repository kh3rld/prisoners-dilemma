package settings

import (
	"fmt"
	"strings"
)

func GetUserSettings() (int, bool) {
	var rounds int
	var detailedSummaries string

	fmt.Print("Enter the number of rounds: ")
	fmt.Scanln(&rounds)

	fmt.Print("Do you want detailed summaries after each round? (y/n): ")
	fmt.Scanln(&detailedSummaries)

	return rounds, strings.ToLower(detailedSummaries) == "y"
}
