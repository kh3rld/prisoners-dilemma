package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/kh3rld/prisoners-dilemma/pkg/common"
)

const (
	broadcastPort     = "8081"
	discoveryMsg      = "DISCOVER"
	broadcastInterval = 5 * time.Second
)

type Server struct {
	IP       string
	Port     string
	UID      string
	Listener net.Listener
	Clients  []net.Conn
}

func NewServer(port string) *Server {
	ip, err := common.GetLocalIP()
	if err != nil {
		log.Fatalf("Error getting local IP: %v", err)
	}
	ln, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	return &Server{IP: ip, Port: port, UID: "", Listener: ln, Clients: make([]net.Conn, 0)}
}

func (s *Server) StartBroadcasting() {
	conn, err := net.Dial("udp", fmt.Sprintf("255.255.255.255:%s", broadcastPort))
	if err != nil {
		log.Fatalf("Error setting up broadcast connection: %v", err)
	}
	defer conn.Close()

	broadcastMessage := fmt.Sprintf("%s:%s:%s", s.IP, s.Port, s.UID)

	for {
		_, err := conn.Write([]byte(broadcastMessage))
		if err != nil {
			log.Printf("Error broadcasting server address: %v", err)
		}
		time.Sleep(broadcastInterval)
	}
}

func (s *Server) ReceiveAction(conn net.Conn) string {
	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error receiving action: %v", err)
		return ""
	}
	return string(buffer[:n])
}

func (s *Server) SendResult(conn net.Conn, result string) {
	_, err := conn.Write([]byte(result))
	if err != nil {
		log.Printf("Error sending result: %v", err)
	}
}

func (s *Server) AcceptConnections() {
	fmt.Println("Server started on", s.IP+":"+s.Port)
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		s.Clients = append(s.Clients, conn)
		if len(s.Clients) >= 2 {
			break
		}
	}
}

func (s *Server) SendName(conn net.Conn, name string) {
	fmt.Fprintf(conn, name+"\n")
}

func (s *Server) ReceiveName(conn net.Conn) string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	return strings.TrimSpace(message)
}

func (s *Server) Close() {
	for _, conn := range s.Clients {
		conn.Close()
	}
	s.Listener.Close()
}
