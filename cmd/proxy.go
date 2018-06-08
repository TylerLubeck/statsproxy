package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerlubeck/statsproxy/pkg"
	"github.com/tylerlubeck/statsproxy/pkg/handlers"
	"os"
	"strings"
)

type HandlerLoaderConfig struct {
	Name string
	Host string
	Port string
	Path string

	Options map[string]interface{}
}

type ProxyConfig struct {
	Host string
	Port string
}

type Config struct {
	Handlers    []HandlerLoaderConfig
	ProxyConfig ProxyConfig
}

var cfgFile string

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Proxies all messages to the configured upstreams",
	Long:  "",
	Run:   RunProxy,
}

func RunProxy(cmd *cobra.Command, args []string) {
	initConfig()

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	s, err := server.NewServer(config.ProxyConfig.Host, config.ProxyConfig.Port)
	if err != nil {
		panic(err)
	}

	h := getHandlers(config.Handlers)

	err = s.Run(h)

	if err != nil {
		panic(err)
	}
}

func getHandlers(handlerConfigs []HandlerLoaderConfig) []server.Handler {

	handlerList := make([]server.Handler, 0, len(handlerConfigs))
	fmt.Println(len(handlerList))

	for _, handlerConfig := range handlerConfigs {
		hc := &server.HandlerConfig{
			Name:    handlerConfig.Name,
			Host:    handlerConfig.Host,
			Port:    handlerConfig.Port,
			Options: handlerConfig.Options,
		}

		if handlerConfig.Path == "echo" {
			h, err := handlers.NewEchoHandler(hc)
			if err != nil {
				fmt.Printf("Failed to set up %s handler: %v\n", hc.Name, err)
				continue
			}

			fmt.Printf("Made handler %s\n", h.Name)
			handlerList = append(handlerList, h)
		} else if handlerConfig.Path == "passthrough" {
			h, err := handlers.NewPassThroughHandler(hc)
			if err != nil {
				fmt.Printf("Failed to set up %s handler: %v\n", hc.Name, err)
				continue
			}

			fmt.Printf("Made handler %s\n", h.Name)
			handlerList = append(handlerList, h)
		} else if strings.HasPrefix(handlerConfig.Path, "plugin:") {
			fmt.Printf("Trying to load a plugin: %s\n", handlerConfig.Path)
		} else {
			fmt.Printf("ERROR: Can't load plugin: %s\n", handlerConfig.Path)
		}
	}

	return handlerList
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config: ", err)
		os.Exit(1)
	}
}

func init() {
	proxyCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	proxyCmd.MarkFlagRequired("config")
}
