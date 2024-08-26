package main

import (
	"fmt"
	"os"

	"github.com/kh3rld/prisoners-dilemma/pkg/network"
	"github.com/kh3rld/prisoners-dilemma/pkg/player"
	"github.com/kh3rld/prisoners-dilemma/pkg/settings"
	"github.com/kh3rld/prisoners-dilemma/pkg/ui"
	"github.com/kh3rld/prisoners-dilemma/pkg/utils"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func DisplayArt() {
	fmt.Println(ui.CenterText(ui.GreenText("===================================")))
	fmt.Println(ui.CenterText(ui.BlueText("  PRISONER'S DILEMMA GAME")))
	fmt.Println(ui.CenterText(ui.GreenText("===================================")))
	fmt.Println()
	fmt.Println(ui.CenterText(ui.GreenText("Developed by Kherld Hussein")))
	fmt.Println(ui.CenterText(ui.GreenText("Special thanks to the Go community")))
	fmt.Println(ui.CenterText(ui.GreenText("and open source contributors.")))
	fmt.Println()
}

func DisplayMenu() {
	ClearScreen()
	DisplayArt()

	for {
		fmt.Println(ui.CenterText("1. Play Locally"))
		fmt.Println(ui.CenterText("2. Play Over Network"))
		fmt.Println(ui.CenterText("3. View Instructions"))
		fmt.Println(ui.CenterText("4. Quit"))

		choice := utils.GetValidatedChoice("Enter your choice (1/2/3/4): ", []string{"1", "2", "3", "4"})

		switch choice {
		case "1":
			StartLocalGame()
		case "2":
			StartNetworkGame()
		case "3":
			ui.DisplayHelp()
		case "4":
			if utils.ConfirmExit() {
				ShowCursor()
				fmt.Println(ui.CenterText("Thank you for playing!"))
				os.Exit(0)
			}
		}
	}

}

func StartLocalGame() {
	utils.ShowProgress("Starting game...")
	player1, player2 := settings.SetPlayers()
	rounds, detailedSummaries := settings.GetUserSettings()

	p1, ok1 := player1.(*player.Player)
	p2, ok2 := player2.(*player.Player)

	if !ok1 || !ok2 {
		fmt.Println("Error: Players must be of type *player.Player")
		return
	}

	p1.SetName(player1.GetName())
	p2.SetName(player2.GetName())

	settings.GameLoop(nil, p1, p2, rounds, detailedSummaries)
}

func StartNetworkGame() {
	fmt.Println(ui.CenterText("1. Host a Game"))
	fmt.Println(ui.CenterText("2. Join a Game"))

	choice := utils.GetValidatedChoice("Enter your choice (1/2): ", []string{"1", "2"})

	switch choice {
	case "1":
		HostGame()
	case "2":
		JoinGame()
	}
}

func HostGame() {
	fmt.Print("Enter your name: ")
	var hostName string
	fmt.Scanln(&hostName)

	server := network.NewServer("8080")
	defer server.Close()

	fmt.Println("Waiting for a client to connect...")
	server.AcceptConnections()

	server.SendName(server.Clients[0], hostName)
	clientName := server.ReceiveName(server.Clients[0])

	fmt.Printf("You are playing against %s\n", clientName)

	player1 := &player.Player{Name: hostName}
	player2 := &player.Player{Name: clientName}

	settings.GameLoop(server, player1, player2, 0, false)
}

func JoinGame() {
	fmt.Print("Enter your name: ")
	var clientName string
	fmt.Scanln(&clientName)

	fmt.Println("Searching for available hosts...")
	hosts, err := network.DiscoverHosts("8080")
	if err != nil {
		fmt.Println("Error discovering hosts:", err)
		return
	}

	if len(hosts) == 0 {
		fmt.Println("No hosts found. Please enter the host IP manually.")
		fmt.Print("Enter the host IP address (or type 'exit' to cancel): ")
		var hostIP string
		fmt.Scanln(&hostIP)

		if hostIP == "exit" {
			fmt.Println("Exiting join game.")
			return
		}

		hosts = append(hosts, hostIP)
	}

	fmt.Println("Available hosts:")
	for i, host := range hosts {
		fmt.Printf("%d. %s\n", i+1, host)
	}

	fmt.Print("Select the host you want to join (enter the number): ")
	var choice int
	fmt.Scanln(&choice)
	if choice < 1 || choice > len(hosts) {
		fmt.Println("Invalid choice. Exiting.")
		return
	}

	hostIP := hosts[choice-1]
	client, err := network.NewClient(hostIP + ":8080")
	if err != nil {
		fmt.Println("Error connecting to host:", err)
		return
	}
	defer client.Close()

	hostName := client.ReceiveName()
	client.SendName(clientName)

	fmt.Printf("You are playing against %s\n", hostName)

	player1 := &player.Player{Name: clientName}
	player2 := &player.Player{Name: hostName}

	settings.GameLoop(client, player1, player2, 0, false)
}
