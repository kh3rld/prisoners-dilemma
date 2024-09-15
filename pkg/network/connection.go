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
	addrInfo, found := cm.peers[peer]
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
	defer s.Close()

	var negotiationMsg NegotiationMessage
	err := json.NewDecoder(s).Decode(&negotiationMsg)
	if err != nil {
		// handle error
		return
	}

	// TODO: Process negotiation message
}

// Handle game state broadcast
func (cm *ConnectionManager) handleGameState(s network.Stream) {
	defer s.Close()

	// Handle game state exchange
	var gameState GameStateMessage
	err := json.NewDecoder(s).Decode(&gameState)
	if err != nil {
		// handle error
		return
	}

	// TODO: Process game state message
}

func (cm *ConnectionManager) Close() error {
	return cm.host.Close()
}
