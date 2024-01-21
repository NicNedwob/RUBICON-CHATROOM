package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var client_map = make(map[string]net.Conn)
var client_list []string

func writeToClient(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
	time.Sleep(100 * time.Millisecond)
}

func readFromClient(conn net.Conn) string {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// handle error
	}
	return msg
}

func writeToAllClients(who_from, msg string) {
	switch msg {
	case "":

	default:
		for _, client := range client_list {
			if who_from != client {
				other_conn := client_map[client]
				msg_sender := strings.TrimSuffix(who_from, "\n")
				msg = strings.TrimSuffix(msg, "\n")
				go writeToClient(other_conn, msg_sender+msg+"\n")
			}
		}
		fmt.Print(strings.TrimSuffix(who_from, "\n") + msg + "\n")

	}

}

func removeClient(user_name string) {
	// find index of client from client list
	var temp []string
	for _, client := range client_list {
		if client != user_name {
			temp = append(temp, client)
		}
	}
	client_list = temp
	delete(client_map, user_name)
	if user_name != "" {
		writeToAllClients(user_name, " has left the chat.")
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// On Board new user and get viable user name
	writeToClient(conn, "Hi New User!\n")
	writeToClient(conn, "People currently in chat: "+strconv.Itoa(len(client_list))+"\n")
	writeToClient(conn, "Please enter a username: \n")
	client_name := ""
	for {
		user_name := readFromClient(conn)
		user_name = strings.TrimSuffix(user_name, "\n")
		if user_name == "Bye" {
			client_name = ""
			break
		}
		_, name_present := client_map[user_name]
		if name_present == false {
			client_map[user_name] = conn
			client_name = user_name
			client_list = append(client_list, user_name)
			writeToClient(conn, "Thank you.\n")
			writeToClient(conn, "Username has been set to: "+user_name+".\n")
			writeToClient(conn, "Here is a list of current users:\n")
			for _, v := range client_list {
				writeToClient(conn, "\t"+v+"\n")
			}
			writeToClient(conn, "Entering chat. Say \"Bye\" or press Ctrl+C to leave.\n")
			writeToClient(conn, "Type and hit enter to send.\n")

			break
		} else {
			writeToClient(conn, "Sorry, that user name is already taken.\n")
			writeToClient(conn, "Here is a list of current users:\n")
			for _, v := range client_list {
				writeToClient(conn, "\t"+v+"\n")
			}
			writeToClient(conn, "Please enter a username: \n")

		}
	}

	// remove the client from the map and list when function ends
	defer removeClient(client_name)

	if client_name != "" {
		writeToAllClients(client_name, " has joined the chat!")
	}

	// read messages from registered user and broadcast to other registered users
	for {
		client_msg := readFromClient(conn)
		if client_name == "" {
			client_msg = "Bye\n"
		}
		if client_msg == "Bye\n" {
			break
		}
		writeToAllClients(client_name, " said: "+client_msg)
	}
}

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
		go handleConnection(conn)
	}
}
