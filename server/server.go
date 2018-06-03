package server

import (
	"fmt"
	"net"
)

type StatsDServer interface {
	Run()
}

type DDServer struct {
	ServerConn *net.UDPConn
}

func (s *DDServer) run(callback func(string, *net.UDPAddr)) {
	fmt.Println("Running...")
	buf := make([]byte, 1024)

	for {
		n, addr, err := s.ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		callback(string(buf[:n]), addr)

	}
}

func (s *DDServer) Echo() {
	s.run(func(msg string, addr *net.UDPAddr) {
		fmt.Println("Received: ", msg, " from ", addr)
	})

}

func (s *DDServer) Proxy() {
	s.run(func(msg string, addr *net.UDPAddr) {
		fmt.Println("Proxying: ", msg, " from ", addr)
		go func() {
			conn, err := net.Dial("udp", "127.0.0.1:6790")
			defer conn.Close()
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(conn, msg)
		}()
	})
}

func NewServer(port string) (*DDServer, error) {
	ServerAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return nil, fmt.Errorf("Bad server addr")
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("Bad server conn")
	}

	return &DDServer{ServerConn: ServerConn}, nil
}
