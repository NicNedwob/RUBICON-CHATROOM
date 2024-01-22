package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/NicNedwob/RUBICON-CHATROOM/client/internal/client"
)

// Global Variables
var srv_msg_struct = client.ServerMsg{}

func main() {
	fmt.Println("Starting Rubicon Chatroom Client")
	scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", "rubiserver:8000")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer client.SendMessageToServer("Bye", conn)

	go srv_msg_struct.ReceiveMessageFromServer(conn)
	if srv_msg_struct.GetMessageContents() == "Please enter a username:\n" {
		for scanner.Scan() {
			msg := scanner.Text()
			client.SendMessageToServer(msg, conn)
		}
	}

	for scanner.Scan() {
		msg := scanner.Text() // Get the current line of text
		client.SendMessageToServer(msg, conn)
	}
}
