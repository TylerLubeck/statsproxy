package commands

import (
	"github.com/spf13/cobra"
	"github.com/tylerlubeck/statsproxy/pkg"
	"github.com/tylerlubeck/statsproxy/pkg/handlers"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echoes all messages sent to it",
	Long:  "",
	Run:   RunEcho,
}

func RunEcho(cmd *cobra.Command, args []string) {
	s, err := server.NewServer("127.0.0.1", "6789")
	if err != nil {
		panic(err)
	}

	h := make([]server.Handler, 1)
	hc := &server.HandlerConfig{
		Name:    "echo",
		Host:    "",
		Port:    "",
		Options: nil,
	}

	h[0], err = handlers.NewEchoHandler(hc)

	err = s.Run(h)

	if err != nil {
		panic(err)
	}
}
