package network

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type ConnectionManager struct {
	connections     map[string]*net.TCPConn
	playerIDMapping map[string]int
	mu              sync.Mutex
}

// NewConnectionManager initializes a new ConnectionManager.
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections:     make(map[string]*net.TCPConn),
		playerIDMapping: make(map[string]int),
		mu:              sync.Mutex{},
	}
}

// AddConnection adds a new connection to the ConnectionManager.
func (cm *ConnectionManager) AddConnection(peerAddr string, conn *net.TCPConn, playerID int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections[peerAddr] = conn
	cm.playerIDMapping[peerAddr] = playerID
	fmt.Printf("Assigned Player %d to connection from %s\n", playerID, peerAddr)
}

// GetConnection retrieves an existing connection by peer address.
func (cm *ConnectionManager) GetConnection(peerAddr string) (*net.TCPConn, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	conn, exists := cm.connections[peerAddr]
	return conn, exists
}

// RemoveConnection removes a connection from the ConnectionManager and closes it.
func (cm *ConnectionManager) RemoveConnection(peerAddr string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if conn, exists := cm.connections[peerAddr]; exists {
		conn.Close()
		delete(cm.connections, peerAddr)
		fmt.Printf("Connection removed: %s\n", peerAddr)
	}
}

// HandleIncomingConnection starts handling data for an incoming connection.
func (cm *ConnectionManager) HandleIncomingConnection(conn *net.TCPConn) {
	fmt.Println("Handling incoming connection from:", conn.RemoteAddr().String())

	// Perform handshake to assign player roles
	playerID := cm.negotiateRole(conn)
	if playerID == -1 {
		fmt.Println("Failed to assign role, closing connection:", conn.RemoteAddr().String())
		conn.Close()
		return
	}
	// Add connection with assigned player ID
	peerAddr := conn.RemoteAddr().String()
	cm.AddConnection(peerAddr, conn, playerID)
	// Handle data in a separate goroutine
	go cm.HandleGameData(conn)
}

func (cm *ConnectionManager) negotiateRole(conn *net.TCPConn) int {
	// Listen for a role request from the connecting peer
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading role request:", err)
		return -1
	}
	request := string(buffer[:n])

	if request == "ROLE_REQUEST" {
		if _, exists := cm.getPlayerWithID(1); !exists {
			// Assign Player 1
			_, err := conn.Write([]byte("ASSIGN_PLAYER_1"))
			if err != nil {
				fmt.Println("Error assigning Player 1:", err)
				return -1
			}
			return 1
		} else if _, exists := cm.getPlayerWithID(2); !exists {
			// Assign Player 2
			_, err := conn.Write([]byte("ASSIGN_PLAYER_2"))
			if err != nil {
				fmt.Println("Error assigning Player 2:", err)
				return -1
			}
			return 2
		} else {
			// No available player slots
			_, err := conn.Write([]byte("NO_AVAILABLE_SLOT"))
			if err != nil {
				fmt.Println("Error sending no slot message:", err)
			}
			return -1
		}
	}

	// Invalid request or no available slots
	fmt.Println("Invalid role request or no available player slots")
	return -1
}

func (cm *ConnectionManager) HandleGameData(conn *net.TCPConn) {
	defer func() {
		fmt.Printf("Closing game connection: %s\n", conn.RemoteAddr().String())
		conn.Close()
	}()

	buffer := make([]byte, 1024)

	for {
		// Read incoming message
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error reading game data from %s: %v\n", conn.RemoteAddr().String(), err)
			return
		}

		message := string(buffer[:n])

		// Process the message based on its type
		switch {
		case strings.HasPrefix(message, "PLAYER_ACTION"):
			action := parseAction(message)
			playerID := cm.getPlayerID(conn)
			processPlayerAction(playerID, action)

		case strings.HasPrefix(message, "GAME_STATE_UPDATE"):
			gameState, err := parseGameState(message)
			if err == nil {
				updateGameState(gameState)
			}

		case strings.HasPrefix(message, "ROUND_OUTCOME"):
			outcome, err := parseOutcome(message)
			if err == nil {
				applyRoundOutcome(outcome, cm)
			}

		default:
			fmt.Printf("Received unknown message type from %s: %s\n", conn.RemoteAddr().String(), message)
		}
	}
}

func (cm *ConnectionManager) getPlayerID(conn *net.TCPConn) int {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	peerAddr := conn.RemoteAddr().String()
	if playerID, exists := cm.playerIDMapping[peerAddr]; exists {
		return playerID
	}

	// Where the player ID is not found
	fmt.Printf("Player ID not found for connection from %s\n", peerAddr)
	return -1 // Invalid player ID
}

func (cm *ConnectionManager) connectToPeer(peer string) {
	addr, err := net.ResolveTCPAddr("tcp", peer)
	if err != nil {
		fmt.Println("Error resolving peer address:", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}

	// Perform the handshake to assign roles
	playerID := cm.performHandshake(conn)
	if playerID == -1 {
		fmt.Println("Handshake failed, closing connection")
		conn.Close()
		return
	}

	// Add connection to ConnectionManager with assigned player ID
	cm.AddConnection(peer, conn, playerID)
	fmt.Printf("Successfully connected to %s as Player %d\n", peer, playerID)
}

func (cm *ConnectionManager) performHandshake(conn *net.TCPConn) int {
	// Send role request
	_, err := conn.Write([]byte("ROLE_REQUEST"))
	if err != nil {
		fmt.Println("Error sending role request:", err)
		return -1
	}

	// Wait for role assignment response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading role assignment:", err)
		return -1
	}
	response := string(buffer[:n])

	// Handle role assignment
	if response == "ASSIGN_PLAYER_1" {
		return 1
	} else if response == "ASSIGN_PLAYER_2" {
		return 2
	}

	// Invalid response
	fmt.Println("Invalid role assignment response:", response)
	return -1
}

func (cm *ConnectionManager) getPlayerWithID(playerID int) (*net.TCPConn, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for peerAddr, id := range cm.playerIDMapping {
		if id == playerID {
			conn, exists := cm.connections[peerAddr]
			if exists {
				return conn, true
			}
		}
	}
	return nil, false
}

func (cm *ConnectionManager) StartGame() {
	// Signal both players to start the game
	for _, conn := range cm.connections {
		_, err := conn.Write([]byte("GAME_START"))
		if err != nil {
			fmt.Println("Error signaling game start:", err)
		}
	}
}
