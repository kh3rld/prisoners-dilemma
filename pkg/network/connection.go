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
	// Handle data in a separate goroutine
	go cm.HandleGameData(conn)
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
			playerID := getPlayerID(conn) // TODO: Define logic to get the player ID
			processPlayerAction(playerID, action)

		case strings.HasPrefix(message, "GAME_STATE_UPDATE"):
			gameState, err := parseGameState(message)
			if err == nil {
				updateGameState(gameState)
			}

		case strings.HasPrefix(message, "ROUND_OUTCOME"):
			outcome, err := parseOutcome(message)
			if err == nil {
				applyRoundOutcome(outcome)
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
