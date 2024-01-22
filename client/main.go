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

	var srv_ipaddrs string
	fmt.Println("Input server IP address:")
	fmt.Scanln(&srv_ipaddrs)
	conn, err := net.Dial("tcp", srv_ipaddrs+":8000") // 172.17.0.2 192.168.1.109

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
