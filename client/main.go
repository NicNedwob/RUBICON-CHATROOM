package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// Structs
type ServerMsg struct {
	server_msg string
	mu         sync.Mutex
}

// Global Variables
var srv_msg_struct = ServerMsg{server_msg: ""}

func SendMessageToServer(msg string, conn net.Conn) {
	switch msg {
	case "Bye":
		msg = msg + "\n"
		conn.Write([]byte(msg))
		conn.Close()
		os.Exit(0)
	case "":
		return
	default:
		msg = msg + "\n"

		conn.Write([]byte(msg))
	}
}

func (srv_msg *ServerMsg) ReceiveMessageFromServer(conn net.Conn) {
	for {
		srv_msg.mu.Lock()

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Server no longer available.")
			fmt.Println("Exiting Application.")
			os.Exit(1)
		}
		srv_msg.server_msg = msg
		fmt.Print(srv_msg.server_msg)
		srv_msg.mu.Unlock()

	}

}

func main() {
	fmt.Println("Starting Rubicon Chatroom Client")
	scanner := bufio.NewScanner(os.Stdin)
	conn, err := net.Dial("tcp", "localhost:8080")
	defer SendMessageToServer("Bye", conn)

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
			SendMessageToServer(msg, conn)
		}
	}

	for scanner.Scan() {
		msg := scanner.Text() // Get the current line of text
		SendMessageToServer(msg, conn)
	}
}
