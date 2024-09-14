package network

import (
	"fmt"
	"net"
)

// manages the sending of discovery requests.
type DiscoveryClient struct {
	conn *net.UDPConn
}

// initializes a new DiscoveryClient.
func NewDiscoveryClient() (*DiscoveryClient, error) {
	addr, err := net.ResolveUDPAddr("udp", discoveryPort)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &DiscoveryClient{conn: conn}, nil
}

// sends a discovery request to find peers.
func (dc *DiscoveryClient) SendDiscoveryRequest() {
	_, err := dc.conn.Write([]byte(broadcastMsg))
	if err != nil {
		fmt.Println("Error sending discovery request:", err)
	}
}

// closes the UDP connection.
func (dc *DiscoveryClient) Close() {
	dc.conn.Close()
}
