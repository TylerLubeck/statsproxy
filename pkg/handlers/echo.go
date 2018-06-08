package handlers

import (
	"fmt"
	"github.com/tylerlubeck/statsproxy/pkg"
)

type EchoHandler struct {
	Name string
}

func NewEchoHandler(config *server.HandlerConfig) (*EchoHandler, error) {
	eh := &EchoHandler{
		Name: config.Name,
	}

	return eh, nil
}

func (eh *EchoHandler) Handle(msg *server.Message) error {
	fmt.Println(msg.ToString())
	return nil
}

func (eh *EchoHandler) Close() error {
	return nil
}

func (eh *EchoHandler) GetName() string {
	return eh.Name
}
