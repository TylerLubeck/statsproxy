package handlers

import (
	"fmt"
	"github.com/tylerlubeck/statsproxy/pkg"
	"net"
)

type PassThroughHandler struct {
	Name       string
	Host       string
	Port       string
	ServerConn *net.UDPConn
}

func NewPassThroughHandler(config *server.HandlerConfig) (*PassThroughHandler, error) {
	var pth PassThroughHandler
	pth.Name = config.Name
	pth.Host = config.Host
	pth.Port = config.Port

	ServerAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(pth.Host, pth.Port))
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve server addres: %v", err)
	}

	ServerConn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial server address: %v", err)
	}

	pth.ServerConn = ServerConn

	return &pth, nil
}

func (pth *PassThroughHandler) Handle(msg *server.Message) error {
	msgString := msg.ToString()
	numBytes, err := fmt.Fprintf(pth.ServerConn, msgString)

	if err != nil {
		return fmt.Errorf("Failed to write message to handler %s: %v", pth.Name, err)
	}

	if numBytes != len(msgString) {
		return fmt.Errorf("Failed to write entire message to handler %s: %d/%d",
			pth.Name,
			numBytes,
			len(msgString))
	}

	return nil
}

func (pth *PassThroughHandler) Close() error {
	err := pth.ServerConn.Close()

	if err != nil {
		return fmt.Errorf("Failed to close server connection for %s: %v", pth.Name, err)
	}

	return nil
}

func (pth *PassThroughHandler) GetName() string {
	return pth.Name
}
