package commands

import (
	"fmt"
	"github.com/TylerLubeck/statsproxy/server"
	"github.com/TylerLubeck/statsproxy/server/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigFileName string

type UpstreamServer struct {
	Address string
	Port    int
	Types   []string
}

type Config struct {
	Address  string
	Port     int
	LogLevel string
}

type ConfigFile struct {
	Upstreams []UpstreamServer
	Config    Config
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run a proxy for statsd events",
	Long:  "",
	Run:   RunProxy,
}

func init() {
	proxyCmd.Flags().StringVarP(&ConfigFileName, "config", "c", "", "The config file")
	proxyCmd.MarkFlagRequired("config")
}

func loadConfigFile() *ConfigFile {
	viper.SetConfigFile(ConfigFileName)
	viper.SetDefault("config.address", "")
	viper.SetDefault("config.port", 6789)
	viper.SetDefault("config.loglevel", "DEBUG")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error from config file: %s \n", err))
	}

	var config ConfigFile
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %s \n", err))
	}

	fmt.Println(config)

	return &config
}

func RunProxy(cmd *cobra.Command, args []string) {
	config := loadConfigFile()

	s, err := server.NewServer(
		utils.FormatServer(config.Config.Address, config.Config.Port))

	if err != nil {
		fmt.Errorf("Failed to build new server: %s \n", err)
	}

	var upstreams []*server.ProxyServer
	upstreams = make([]*server.ProxyServer, len(config.Upstreams))
	for idx, upstream := range config.Upstreams {
		upstreamServer, err := server.NewProxy(
			utils.FormatServer(upstream.Address, upstream.Port))

		if err != nil {
			fmt.Println(err)
			continue
		}

		upstreams[idx] = upstreamServer
	}

	s.Proxy(upstreams)
}
