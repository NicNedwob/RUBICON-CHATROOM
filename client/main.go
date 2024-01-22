package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/NicNedwob/RUBICON-CHATROOM/client/internal/client"
)

// server message struct (contains mutex to make sure messages not overriden)
var srv_msg_struct = client.ServerMsg{}

func main() {
	// print Terminal Heart
	fmt.Println("RubiChat Client")

	// ask user for server ip address
	var srv_ipaddrs string
	fmt.Println("Input server IP address:")
	fmt.Scanln(&srv_ipaddrs)
	conn, err := net.Dial("tcp", srv_ipaddrs+":8000") // dial server on port 8000

	// if cannot connect show error and exit with fail
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// no matter how main exits, client will send Bye message to server to gracefully exit chat and close connection.
	defer client.SendMessageToServer("Bye", conn)

	// load messages from server using goroutine into server struct to be displayed a
	go srv_msg_struct.ReceiveMessageFromServer(conn)

	// Starts the loop to scan for messages and send them to receiver
	scanner := bufio.NewScanner(os.Stdin) // scanner for using input during messaging
	for scanner.Scan() {
		msg := scanner.Text()                 // Get the current line of text
		client.SendMessageToServer(msg, conn) // Send that text to the server
	}
}
