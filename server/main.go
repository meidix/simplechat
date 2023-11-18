package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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


	gm := make(chan string)

	go messegeDispatcher(users, gm)

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("There was a problem in establishing the connection")
			os.Exit(1)
		}
		rbuff := make([]byte, 1024)
		wbuff := []byte("Hello There, What is your name: ")
		conn.Write(wbuff)
		_, err = conn.Read(rbuff)
		if err != nil {
			fmt.Println("There was a problem reading response")
			os.Exit(1)
		}
		setupChat(rbuff, conn, users, gm)
		fmt.Println("user connected", strings.Trim(string(rbuff), "\x00"), "connected")

	}
}

func setupChat(buff []byte, conn *net.TCPConn, u map[string]net.TCPConn, ch chan string) {
	key := strings.Trim(string(buff), "\x00")
	u[key] = *conn
	go messegeListener(key, *conn, ch)
	conn.Write([]byte("Welcome to This Chat"))
}


func messegeListener(u string, conn net.TCPConn, ch chan string) {
	for {
		rbuff := make([]byte, 1024)
		_, err := conn.Read(rbuff); if err != nil {
			fmt.Println("There was a problem receiving messege")
			os.Exit(1)
		}
		ch <- strings.Trim(string(rbuff), "\x00")
	}
}


func messegeDispatcher(u map[string]net.TCPConn, ch chan string) {
	for m := range ch {
		for key, conn := range u {
			_, err := conn.Write([]byte(m)); if err != nil {
				fmt.Println("there was a problem sending messeges")
				delete(u, key)
			}
		}
	}
}
