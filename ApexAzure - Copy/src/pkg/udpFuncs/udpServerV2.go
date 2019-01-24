package udpFuncs

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// Server defines the minimum contract our TCP and UDP server implementations must satisfy.
type Server interface {
	Run() error
	Close() error
}

// NewServer creates a new Server using given protocol and addr.
func NewServer(protocol, addr string) (Server, error) {
  
	switch strings.ToLower(protocol) {
	case "udp":
		return &UDPServer{
			addr: addr,
		}, nil
	}
	return nil, errors.New("Invalid protocol given")
}

// UDPServer holds the necessary structure for our UDP server.
type UDPServer struct {
	addr   string
	server *net.UDPConn
}

// Run starts the UDP server.
func (u *UDPServer) Run() (err error) {

	laddr, err := net.ResolveUDPAddr("udp", u.addr)
	if err != nil {
		return errors.New("could not resolve UDP addr")
	}

	u.server, err = net.ListenUDP("udp", laddr)
	
	if err != nil {
		return errors.New("could not listen on UDP")
	}

	return u.handleConnections()
}

func (u *UDPServer) handleConnections() error {

	var err error

	for {
		buf := make([]byte, 2048)
		n, conn, err := u.server.ReadFromUDP(buf)

		if err != nil {
			log.Println(err)
			break
		}

		if conn == nil {
			continue
		}

		go u.handleConnection(conn, buf[:n])
	}
	return err
}

func (u *UDPServer) handleConnection(addr *net.UDPAddr, cmd []byte) {
	u.server.WriteToUDP([]byte(fmt.Sprintf("Request recieved: %s", cmd)), addr)
	fmt.Println("UDP Working")
	fmt.Println()
}

// Close ensures that the UDPServer is shut down gracefully.
func (u *UDPServer) Close() error {
	return u.server.Close()
}