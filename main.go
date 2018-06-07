package main

import (
	"github.com/TylerLubeck/statsproxy/cmd"
)

func main() {
	/*
		s, err := server.NewServer("127.0.0.1", "6789")
		if err != nil {
			panic(err)
		}

		h := make([]server.Handler, 1)

		hc := &server.HandlerConfig{
			Name:    "test",
			Host:    "127.0.0.1",
			Port:    "6789",
			Options: nil,
		}
		h[0], err = server.NewHandler(hc)

		err = s.Run(h)

		if err != nil {
			panic(err)
		}
	*/
	commands.Execute()
}
