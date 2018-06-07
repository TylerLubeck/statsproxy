package server

import (
	"fmt"
	"net"
)

type Server struct {
	ServerConn *net.UDPConn
}

func NewServer(host, port string) (*Server, error) {

	ServerAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve server addres: %v", err)
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to listen on server address: %v", err)
	}

	return &Server{ServerConn: ServerConn}, nil
}

func (s *Server) Run(handlers []Handler) error {
	buf := make([]byte, 1024)
	defer s.ServerConn.Close()

	for {
		n, _, err := s.ServerConn.ReadFromUDP(buf)
		if err != nil {
			// TODO: what?
		}

		m := NewMessage(string(buf[:n]))

		for _, handler := range handlers {
			go func() {
				handlerErr := handler.Handle(m)
				if handlerErr != nil {
					// TODO: log?
				}
			}()
		}
	}

	return nil
}
