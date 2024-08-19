package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	Address  string
	Listener net.Listener
	Clients  []net.Conn
}

func NewServer(port string) *Server {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	return &Server{Listener: ln, Clients: make([]net.Conn, 0)}
	// 	// return &Server{Address: port}
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
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	s.Listener = listener
	fmt.Println("Server started on", s.Address)

	for {
		conn, err := listener.Accept()
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
	return message[:len(message)-1]
}

func (s *Server) Close() {
	for _, conn := range s.Clients {
		conn.Close()
	}
	s.Listener.Close()
}
