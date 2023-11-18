package main

import (
	"fmt"
	"meidix/chatserv/chat"
	"meidix/chatserv/connection"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("There was a problem resolving the address")
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("There was a problem in running the server")
		os.Exit(1)
	}

	fmt.Println("The server is listenning on", l.Addr().String())

	users := map[string]net.TCPConn{}

	gm := make(chan chat.Messege)

	go chat.MessegeDispatcher(users, gm)

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("There was a problem in establishing the connection")
			os.Exit(1)
		}

		err = connection.SendTxtMsg(conn, "Hello There, What is you name: "); if err != nil {
			fmt.Println("There was a problem sending a text messege", err)
		}

		user, err := connection.RecvTxtMsg(conn); if err != nil {
			fmt.Println("There was a problem receiving the messege")
			os.Exit(1)
		}
		chat.SetupChat(user, conn, users, gm)
		fmt.Println("user connected", user, "connected")

	}
}
