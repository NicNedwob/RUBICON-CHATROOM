package main

import (
	"fmt"
	"net"

	"github.com/NicNedwob/RUBICON-CHATROOM/server/internal/server"
)

// Global Variables

func main() {
	fmt.Println("Starting Rubicon Chatroom Server")
	ln, err := net.Listen("tcp", "137.125.36.35:5000") // localhost
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}

		// create new go routine for each client
		go server.HandleConnection(conn)
	}
}
