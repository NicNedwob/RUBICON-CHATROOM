package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// non-exported
var client_map = make(map[string]net.Conn) // maps user names to connections
var client_list []string                   // stores list of current active user names

func writeToClient(conn net.Conn, msg string) {
	conn.Write([]byte(msg))            // write message to client over conn
	time.Sleep(100 * time.Millisecond) //wait 100ms to ensure client received message before sending next message
}

func readFromClient(conn net.Conn) string {
	// read data into message and return string including \n
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	return msg
}

func writeToAllClients(who_from, msg string) {
	switch msg {
	case "":
		// do nothing if message was empty (client just hit enter without typing)
	default:
		// usual case of client types somethng and sends
		// iterate through client list
		for _, client := range client_list {

			// if current client in list is not send, send the message
			if who_from != client {

				other_conn := client_map[client] // get connection from map

				// make sure no "\n" are present in message
				msg_sender := strings.TrimSuffix(who_from, "\n")
				msg = strings.TrimSuffix(msg, "\n")

				// start a go routine to write to client so all clients receive message at same time
				go writeToClient(other_conn, msg_sender+msg+"\n")
			}
		}

		// print the message to server console
		fmt.Print(strings.TrimSuffix(who_from, "\n") + msg + "\n")

	}

}

func removeClient(user_name string) {
	// remove username from list using append method
	var temp []string
	for _, client := range client_list {
		if client != user_name {
			temp = append(temp, client)
		}
	}
	client_list = temp // set client list to list without user_name

	// remove client from map
	delete(client_map, user_name)

	// broadcast that client has left the chat.
	if user_name != "" {
		writeToAllClients(user_name, " has left the chat.")
	}
}

// exported
func HandleConnection(conn net.Conn) {
	// defer closing of connection till end of function (gaurantees graceful closure)
	defer conn.Close()

	// On Board new user and get viable user name
	writeToClient(conn, "Hi New User!\n")
	writeToClient(conn, "People currently in chat: "+strconv.Itoa(len(client_list))+"\n")
	writeToClient(conn, "Please enter a username: \n")

	// user name selection loop (will continue till receives buy or valid name)
	client_name := "" // set an empty client name to be used later
	for {
		// client user name attemps are always first thing to be transmitted
		user_name := readFromClient(conn) // store user name
		user_name = strings.TrimSuffix(user_name, "\n")

		// if client exited prematurely user_name would be Bye sp break loop
		if user_name == "Bye" {
			client_name = ""
			break
		}

		// check if user name is already present in client map
		_, name_present := client_map[user_name]
		if name_present == false {

			// if username is unique -> register user
			client_map[user_name] = conn
			client_name = user_name
			client_list = append(client_list, user_name)
			writeToClient(conn, "Thank you.\n")
			writeToClient(conn, "Username has been set to: "+user_name+".\n")
			writeToClient(conn, "Here is a list of current users:\n")

			// provide a list of current chat users
			for _, v := range client_list {
				writeToClient(conn, "\t"+v+"\n")
			}

			// provide instructions on how to use chat.
			writeToClient(conn, "Entering chat. Say \"Bye\" or press Ctrl+C to leave.\n")
			writeToClient(conn, "Type and hit enter to send.\n")

			break //proceed to next loop for message broadcasting

		} else {

			// if user name is not unique ask for another name
			writeToClient(conn, "Sorry, that user name is already taken.\n")
			writeToClient(conn, "Here is a list of current users:\n")

			// provide a list of current users so they can find a unique name
			for _, v := range client_list {
				writeToClient(conn, "\t"+v+"\n")
			}
			writeToClient(conn, "Please enter a username: \n")

		}
	}

	// remove the client from the map and list when function ends.
	// is added on top of defer stack so will run before connection is closed by first defer
	defer removeClient(client_name)

	// if client did not leave prematurely (no username selected)
	if client_name != "" {

		// tell clients new client has joined the chat
		writeToAllClients(client_name, " has joined the chat!")
	}

	// read messages from registered user and broadcast to other registered users
	for {
		// read incoming messages from client using connection conn
		client_msg := readFromClient(conn)

		// these if statements are a bit hacky but work
		// ensures that if client left or left before selecting user name that
		// the server does not spam chat with empty messages
		if client_name == "" {
			client_msg = "Bye\n"
			break
		}
		if client_msg == "Bye\n" {
			break
		}
		if client_msg == "" {
			break
		}
		writeToAllClients(client_name, " said: "+client_msg)
	}

	// function is closed if last loop broken and deferred are executed to remove client and close connection
}
