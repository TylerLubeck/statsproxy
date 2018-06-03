package commands

import (
	"fmt"
	"github.com/TylerLubeck/statsproxy/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigFile string

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run a proxy for statsd events",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(ConfigFile)
		s, _ := server.NewServer(fmt.Sprintf(":%s", Port))
		s.Proxy()
	},
}

func init() {
	//proxyCmd.Flags().StringVarP(&ConfigFile, "config", "c", "", "The config file")
	//proxyCmd.MarkFlagRequired("config")
}
