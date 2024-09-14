package network

import (
	"fmt"
	"net"
	"time"
)

const (
	discoveryPort = ":9999" // Port for discovery messages
	broadcastMsg  = "DISCOVERY_REQUEST"
	responseMsg   = "DISCOVERY_RESPONSE"
)

type DiscoveryService struct {
	conn  *net.UDPConn
	peers []string
}

func NewDiscoveryService() (*DiscoveryService, error) {
	addr, err := net.ResolveUDPAddr("udp", discoveryPort)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &DiscoveryService{conn: conn, peers: make([]string, 0)}, nil
}

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
			ds.peers = append(ds.peers, addr.String())
		}
	}
}

func (ds *DiscoveryService) GetDiscoveredPeers() []string {
	return ds.peers
}
