package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "statsproxy",
	Short: "statsproxy is a statsd server and proxy",
	Long:  "",
}

func Execute() {

	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(proxyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
