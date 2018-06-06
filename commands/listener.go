package commands

import (
	"fmt"
	"github.com/TylerLubeck/statsproxy/server"
	"github.com/spf13/cobra"
)

var Port string

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for UDP messages and print them to the screen",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := server.NewServer(fmt.Sprintf(":%s", Port))
		if err != nil {
			panic(err)
		}
		s.Echo()
	},
}

func init() {
	listenCmd.Flags().StringVarP(&Port, "port", "p", "", "The port to listen on")
	listenCmd.MarkFlagRequired("config")
}
