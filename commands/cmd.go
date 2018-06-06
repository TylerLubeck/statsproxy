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
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {

	rootCmd.AddCommand(listenCmd)
	rootCmd.AddCommand(proxyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
