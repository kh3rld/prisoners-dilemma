package network

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type ConnectionManager struct {
	host  host.Host
	peers map[peer.ID]peer.AddrInfo
	mu    sync.Mutex
}

type NegotiationMessage struct {
	PlayerID    string `json:"player_id"`
	MessageType string `json:"message_type"`
	Content     string `json:"content"`
}

type GameStateMessage struct {
	PlayerID  string `json:"player_id"`
	GameState string `json:"game_state"`
	Round     int    `json:"round"`
}

// NewConnectionManager initializes a new ConnectionManager.
func NewConnectionManager() (*ConnectionManager, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}

	cm := &ConnectionManager{
		host:  h,
		peers: make(map[peer.ID]peer.AddrInfo),
	}

	return cm, nil
}

// AddConnection adds a new connection to the ConnectionManager.
func (cm *ConnectionManager) AddConnection(addrInfo peer.AddrInfo) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.peers[addrInfo.ID] = addrInfo
}

func (cm *ConnectionManager) connectToPeer(peer peer.ID) error {
	cm.mu.Lock()
	addrInfo, found := cm.peers[peer]
	cm.mu.Unlock()
	if !found {
		return fmt.Errorf("peer not found")
	}
	return cm.host.Connect(context.Background(), addrInfo)
}

func (cm *ConnectionManager) RegisterProtocols() {
	cm.host.SetStreamHandler(protocol.ID("/pdgame/negotiation/1.0.0"), cm.handleNegotiation)
	cm.host.SetStreamHandler(protocol.ID("/pdgame/gamestate/1.0.0"), cm.handleGameState)
}

func (cm *ConnectionManager) handleNegotiation(s network.Stream) {
	go func() {
		defer s.Close()

		var negotiationMsg NegotiationMessage
		err := json.NewDecoder(s).Decode(&negotiationMsg)
		if err != nil {
			fmt.Printf("Failed to decode negotiation message: %v\n", err)
			return
		}

		// Process negotiation message
		fmt.Printf("Received negotiation message: %+v\n", negotiationMsg)

		// Send a response
		response := NegotiationMessage{
			// TODO: Implement
		}
		if err := json.NewEncoder(s).Encode(response); err != nil {
			fmt.Printf("Failed to send negotiation response: %v\n", err)
		}
	}()
}

// Handle game state broadcast
func (cm *ConnectionManager) handleGameState(s network.Stream) {
	defer s.Close()

	// Handle game state exchange
	var gameState GameStateMessage
	err := json.NewDecoder(s).Decode(&gameState)
	if err != nil {
		fmt.Printf("Failed to decode game state message: %v\n", err)
		return
	}
	// Process game state message
	fmt.Printf("Received game state: %+v\n", gameState)
}

func (cm *ConnectionManager) RegisterPeerDisconnectHandler() {
	cm.host.Network().Notify(&network.NotifyBundle{
		DisconnectedF: func(n network.Network, conn network.Conn) {
			peerID := conn.RemotePeer()
			fmt.Printf("Peer %s disconnected\n", peerID)

			cm.mu.Lock()
			delete(cm.peers, peerID)
			cm.mu.Unlock()
		},
	})
}

func (cm *ConnectionManager) reconnectToPeer(peerID peer.ID) error {
	cm.mu.Lock()
	addrInfo, found := cm.peers[peerID]
	cm.mu.Unlock()

	if !found {
		return fmt.Errorf("peer not found for reconnection")
	}

	fmt.Printf("Attempting to reconnect to peer %s\n", peerID)
	return cm.host.Connect(context.Background(), addrInfo)
}

func (cm *ConnectionManager) Close() error {
	return cm.host.Close()
}
