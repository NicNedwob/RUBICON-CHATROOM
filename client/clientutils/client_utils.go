package clientutils

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
