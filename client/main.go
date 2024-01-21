package main

import (
	"bufio"
	"client/clientutils"
	"fmt"
	"net"
	"os"
)

// Global Variables
var srv_msg_struct = clientutils.ServerMsg{server_msg: ""}

func main() {
	fmt.Println("Starting Rubicon Chatroom Client")
	scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", "localhost:8080")
	defer clientutils.SendMessageToServer("Bye", conn)

	if err != nil {
		// handle error
	}
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')

	// Set username
	go srv_msg_struct.ReceiveMessageFromServer(conn)
	if srv_msg_struct.server_msg == "Please enter a username:\n" {
		for scanner.Scan() {
			msg := scanner.Text()
			clientutils.SendMessageToServer(msg, conn)
		}
	}

	for scanner.Scan() {
		msg := scanner.Text() // Get the current line of text
		clientutils.SendMessageToServer(msg, conn)
	}
}
