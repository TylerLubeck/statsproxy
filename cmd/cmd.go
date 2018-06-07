package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tylerlubeck/statsproxy/pkg"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "statsproxy",
	Short: "statsproxy is a statsd server and proxy",
	Long:  "",
}

func Execute() {

	rootCmd.AddCommand(echoCmd)
	//rootCmd.AddCommand(proxyCmd)

	m := server.NewMessage("hello there")
	fmt.Println(m.ToString())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
