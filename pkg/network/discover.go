package network

import (
	"fmt"
	"net"
	"time"
)

func DiscoverHosts(port string) ([]string, error) {
	var hosts []string

	conn, err := net.Dial("udp", "255.255.255.255:"+broadcastPort)
	if err != nil {
		return nil, fmt.Errorf("error creating UDP connection: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(discoveryMsg))
	if err != nil {
		return nil, fmt.Errorf("error sending discovery message: %v", err)
	}

	listen, err := net.ListenPacket("udp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("error listening for responses: %v", err)
	}
	defer listen.Close()

	listen.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 1024)
	for {
		n, addr, err := listen.ReadFrom(buffer)
		if err != nil {
			break
		}

		message := string(buffer[:n])
		if message == discoveryMsg {
			hosts = append(hosts, addr.String())
		}
	}

	return hosts, nil
}
