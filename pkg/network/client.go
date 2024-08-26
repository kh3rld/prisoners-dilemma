package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Client struct {
	conn net.Conn
	Name string
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	return &Client{conn: conn}, nil
}

func ListenForHosts(port string) ([]string, error) {
	conn, err := net.ListenPacket("udp4", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("error listening for hosts: %v", err)
	}
	defer conn.Close()

	hosts := make(map[string]struct{})
	buffer := make([]byte, 1024)

	for {
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, _, err := conn.ReadFrom(buffer)
		if err != nil {
			break
		}

		host := strings.TrimSpace(string(buffer[:n]))
		if _, exists := hosts[host]; !exists {
			fmt.Println("Discovered host:", host)
			hosts[host] = struct{}{}
		}
	}

	hostList := make([]string, 0, len(hosts))
	for host := range hosts {
		hostList = append(hostList, host)
	}

	return hostList, nil
}

func ChooseHost(hosts []string) string {
	if len(hosts) == 0 {
		fmt.Println("No hosts found. Try again later.")
		return ""
	}

	fmt.Println("Available Hosts:")
	for i, host := range hosts {
		fmt.Printf("%d. %s\n", i+1, host)
	}

	var choice int
	for {
		fmt.Print("Select a host to connect to: ")
		_, err := fmt.Scan(&choice)
		if err == nil && choice > 0 && choice <= len(hosts) {
			break
		}
		fmt.Println("Invalid choice. Please try again.")
	}

	return hosts[choice-1]
}

func (c *Client) SendAction(action string) {
	_, err := c.conn.Write([]byte(action))
	if err != nil {
		log.Printf("Error sending action: %v", err)
	}
}

func (c *Client) ReceiveResult() string {
	buffer := make([]byte, 256)
	n, err := c.conn.Read(buffer)
	if err != nil {
		log.Printf("Error receiving result: %v", err)
		return ""
	}
	return string(buffer[:n])
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SendName(name string) {
	c.Name = name
	fmt.Fprintf(c.conn, name+"\n")
}

func (c *Client) ReceiveName() string {
	message, _ := bufio.NewReader(c.conn).ReadString('\n')
	return strings.TrimSpace(message)
}
