package commands

import (
	"fmt"
	"github.com/TylerLubeck/statsproxy/server"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for UDP messages and print them to the screen",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := server.NewServer(fmt.Sprintf(":%s", Port))
		s.Echo()
	},
}

func init() {
}
