package network

import (
	"fmt"
	"net"
)

type DiscoveryClient struct {
	conn *net.UDPConn
}

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

func (dc *DiscoveryClient) SendDiscoveryRequest() {
	_, err := dc.conn.Write([]byte(broadcastMsg))
	if err != nil {
		fmt.Println("Error sending discovery request:", err)
	}
}

func (dc *DiscoveryClient) Close() {
	dc.conn.Close()
}
