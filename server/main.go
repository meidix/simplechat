package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("There was a problem in running the server")
		os.Exit(1)
	}

	fmt.Println("The server is listenning on", l.Addr().String())

	users := map[string]net.TCPConn{}

	for i := 0; i < 2; i++{
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
		users[strings.Trim(string(rbuff), "\x00")] = *conn
		fmt.Println("user connected", strings.Trim(string(rbuff), "\x00"), "connected")

	}

	fmt.Println(users)

	g := []byte("You can now start chatting")

	gm := make(chan string)

	go messegeDispatcher(users, gm)

	for key, conn := range users {
		_, err := conn.Write(g); if err != nil {
			fmt.Println("There was a problem sending a messege")
		}
		go messegeListener(key, conn, gm)
	}

	end := ""
	fmt.Scan(&end)
}


func messegeListener(u string, conn net.TCPConn, ch chan string) {
	for {
		rbuff := make([]byte, 1024)
		_, err := conn.Read(rbuff); if err != nil {
			fmt.Println("There was a problem receiving messege")
		}
		ch <- strings.Trim(string(rbuff), "\x00")
	}
}


func messegeDispatcher(u map[string]net.TCPConn, ch chan string) {
	for m := range ch {
		for _, conn := range u {
			_, err := conn.Write([]byte(m)); if err != nil {
				fmt.Println("there was a problem sending messeges")
			}
		}
	}
}
