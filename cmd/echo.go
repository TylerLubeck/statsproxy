package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tylerlubeck/statsproxy/pkg"
	"github.com/tylerlubeck/statsproxy/pkg/handlers"
)

var echoHost = "127.0.0.1"
var echoPort = "6789"

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echoes all messages sent to it",
	Long:  "",
	Run:   RunEcho,
}

func RunEcho(cmd *cobra.Command, args []string) {
	s, err := server.NewServer(echoHost, echoPort)
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

	fmt.Printf("Listening on %s...\n", s.ServerAddr.String())

	err = s.Run(h)

	if err != nil {
		panic(err)
	}
}

func init() {
	echoCmd.Flags().StringVar(&echoHost, "host", echoHost, "host to listen on")
	echoCmd.Flags().StringVar(&echoPort, "port", echoPort, "port to listen on")
}
