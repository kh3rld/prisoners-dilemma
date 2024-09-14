package network

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	discoveryPort = ":9999"
	broadcastMsg  = "DISCOVERY_REQUEST"
	responseMsg   = "DISCOVERY_RESPONSE"
)

type DiscoveryService struct {
	conn  *net.UDPConn
	cm    *ConnectionManager
	mu    sync.Mutex
	peers []string
}

func NewDiscoveryService(cm *ConnectionManager) (*DiscoveryService, error) {
	addr, err := net.ResolveUDPAddr("udp", discoveryPort)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &DiscoveryService{
		conn:  conn,
		cm:    cm,
		peers: make([]string, 0),
	}, nil
}

// sends discovery requests periodically.
func (ds *DiscoveryService) BroadcastDiscovery() {
	addr, _ := net.ResolveUDPAddr("udp", discoveryPort)
	for {
		_, err := ds.conn.WriteToUDP([]byte(broadcastMsg), addr)
		if err != nil {
			fmt.Println("Error broadcasting discovery message:", err)
		}
		time.Sleep(5 * time.Second) // Broadcast every 5 seconds
	}
}

// listens for discovery responses and stores peer addresses.
func (ds *DiscoveryService) ListenForPeers() {
	buffer := make([]byte, 1024)
	for {
		n, addr, err := ds.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}

		msg := string(buffer[:n])
		if msg == responseMsg {
			fmt.Printf("Received response from %s\n", addr.String())
			ds.mu.Lock()
			ds.peers = append(ds.peers, addr.String())
			ds.mu.Unlock()
		}
	}
}

// returns a slice of discovered peers.
func (ds *DiscoveryService) GetDiscoveredPeers() []string {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	return ds.peers
}

// starts the discovery and connection process
func (ds *DiscoveryService) DiscoverAndConnect() {
	go ds.BroadcastDiscovery()
	go ds.ListenForPeers()

	time.Sleep(10 * time.Second) // Wait for discovery

	peers := ds.GetDiscoveredPeers()
	fmt.Printf("Discovered peers: %v\n", peers)

	// Initiate TCP connections to discovered peers
	for _, peer := range peers {
		go ds.connectToPeer(peer)
	}
}

func (ds *DiscoveryService) connectToPeer(peer string) {
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

	// Add connection to ConnectionManager
	ds.cm.AddConnection(peer, conn)
	fmt.Printf("Successfully connected to %s\n", peer)
}
