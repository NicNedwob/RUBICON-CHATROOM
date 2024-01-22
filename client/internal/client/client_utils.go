package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// Structs
type ServerMsg struct {
	server_msg string     // message string
	mu         sync.Mutex // mutex to ensure sync
}

func SendMessageToServer(msg string, conn net.Conn) {
	switch msg {
	case "Bye":
		// exit message that closes connection
		msg = msg + "\n"
		conn.Write([]byte(msg))
		conn.Close()
		os.Exit(0)
	case "":
		// does not anything in case of empty message
		return
	default:
		// usual case, just write the message
		msg = msg + "\n"

		conn.Write([]byte(msg))
	}
}

func (srv_msg *ServerMsg) GetMessageContents() string {
	// just a get method
	return srv_msg.server_msg
}

func (srv_msg *ServerMsg) ReceiveMessageFromServer(conn net.Conn) {
	// is
	// constantly blocks waiting for messagaes from receiver
	// this is run in a go routine so rest of program can continue
	for {
		// just making sure it can't be edited at same time
		srv_msg.mu.Lock()

		// wait for message
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// if there was an error (i.e. the conenctio was closed), exit
			fmt.Println("Server closed. Exiting application.")
			os.Exit(1)
		}

		// set value in server message struct
		srv_msg.server_msg = msg

		// print the msg to console
		fmt.Print(srv_msg.server_msg)
		srv_msg.mu.Unlock()
	}

}
