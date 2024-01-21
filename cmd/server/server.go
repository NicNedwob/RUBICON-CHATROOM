package main

import (
	"fmt"
	"net"

	"github.com/NicNedwob/RUBICON-CHATROOM/internal/server"
)

// Global Variables

func main() {
	fmt.Println("Starting Rubicon Chatroom Server")
	ln, err := net.Listen("tcp", "localhost:8080")
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
