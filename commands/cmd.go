package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Port string

var rootCmd = &cobra.Command{
	Use:   "statsproxy",
	Short: "statsproxy is a statsd server and proxy",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "", "The port to listen on")
	rootCmd.MarkFlagRequired("port")

	rootCmd.AddCommand(listenCmd)
	rootCmd.AddCommand(proxyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
