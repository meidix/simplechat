package main

import (
	"fmt"
	"meidix/chatclient/chatclient"
	"meidix/chatclient/connection"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("there was a problem in resolving the addr")
	}

	conn, err := net.DialTCP("tcp", nil, addr); if err != nil {
		fmt.Println("Could not connect to server, error is :", err)
		os.Exit(1)
	}

	resp, err := connection.RecvTxtMsg(conn)
	if err != nil {
		fmt.Println("There was a problem receiving the messege")
		os.Exit(1)
	}
	fmt.Println(resp)

	user := ""
	fmt.Scan(&user)
	err = connection.SendTxtMsg(conn, user)
	if err != nil {
		fmt.Println("There was a problem sending a messege")
	}

	ch := make(chan int)

	go chatclient.EstablishReceiver(*conn, ch)
	go chatclient.EstablishSender(*conn, ch)

	<- ch

}
