package server

import (
	"fmt"
	"github.com/TylerLubeck/statsproxy/server/datatypes"
	"net"
	"strings"
)

type StatsDServer struct {
	ServerConn *net.UDPConn
}

type ProxyServer struct {
	Address    string
	Connection *net.UDPConn
	MsgChan    <-chan string
}

func (s *StatsDServer) run(callback func(string, *net.UDPAddr)) {
	fmt.Println("Running...")
	buf := make([]byte, 1024)
	defer s.ServerConn.Close()

	for {
		n, addr, err := s.ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(n, string(buf[:n]))
		callback(string(buf[:n]), addr)
	}
}

func (s *StatsDServer) Echo() {
	s.run(func(msg string, addr *net.UDPAddr) {
		fmt.Println("Received: ", msg, " from ", addr)
	})

}

func checkPacketForUpstream(p datatypes.DataType) bool {
	return true
}

func (s *StatsDServer) Proxy(upstreams []*ProxyServer) {

	/*
		for _, server := range upstreams {
			defer server.Connection.Close()
		}
	*/

	s.run(func(msg string, addr *net.UDPAddr) {
		go func() {
			fmt.Println("got this message: ", msg)
			for _, m := range strings.Split(msg, "\n") {
				if len(strings.TrimSpace(m)) == 0 {
					continue
				}
				for _, server := range upstreams {
					packet := datatypes.ParseDataPacket(m)
					fmt.Println(packet.ToString())

					/*
						if !checkPacketForUpstream(packet) {
							continue
						}
					*/

					fmt.Println("Proxying: ", m, " from ", addr)
					fmt.Fprintf(server.Connection, m)
				}
			}
		}()
	})
}

func NewServer(host string) (*StatsDServer, error) {
	fmt.Printf("Starting on port %s...\n", host)
	ServerAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return nil, fmt.Errorf("Bad server addr")
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("Bad server conn")
	}

	return &StatsDServer{ServerConn: ServerConn}, nil
}

func NewProxy(host string) (*ProxyServer, error) {
	ServerAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return nil, fmt.Errorf("Bad server addr")
	}

	ServerConn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect")
	}

	var msgChan chan string

	msgChan = make(chan string)

	server := &ProxyServer{
		Address:    host,
		Connection: ServerConn,
		MsgChan:    msgChan,
	}

	return server, nil
}
