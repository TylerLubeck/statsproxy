package main

import (
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":6789")
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	fmt.Println("Listening on: :6789...")

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received: ", string(buf[:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
