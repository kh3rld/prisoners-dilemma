package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(serverAddr string) *Client {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	return &Client{conn: conn}
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
	fmt.Fprintf(c.conn, name+"\n")
}

func (c *Client) ReceiveName() string {
	message, _ := bufio.NewReader(c.conn).ReadString('\n')
	return message[:len(message)-1]
}
