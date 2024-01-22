package main

import (
	"fmt"
	"net"

	"github.com/NicNedwob/RUBICON-CHATROOM/server/internal/server"
)

// Global Variables

func main() {
	fmt.Println("RubiChat Server")

	// listen on everything (within container)
	ln, err := net.Listen("tcp", "0.0.0.0:8000") // localhost
	if err != nil {
		fmt.Println(err.Error())
	}

	// main loop accepts new connections and makes a new go routine to handle the client till they leave
	for {

		// try accept connection
		conn, err := ln.Accept()
		if err != nil {
			//print if there is error
			fmt.Println(err.Error())
		} else {
			// create new go routine for per client if successful
			go server.HandleConnection(conn)
		}

	}
}
