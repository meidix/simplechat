package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	buff := make([]byte, 1024)
	_ , err = conn.Read(buff); if err != nil {
		fmt.Println("There was a problem receiving the messege")
		os.Exit(1)
	}
	fmt.Println(strings.Trim(string(buff), "\x00"))

	user := ""
	fmt.Scan(&user)
	buff = []byte(user)
	_, err = conn.Write(buff); if err != nil {
		fmt.Println("There was a problem sending a messege")
	}

	ch := make(chan int)

	go establishReceiver(*conn, ch)
	go establishSender(*conn, ch)

	<- ch

}

func establishReceiver(conn net.TCPConn, c chan int) {
	for {
		rbuff := make([]byte, 1024)
		_, err := conn.Read(rbuff); if err != nil {
			fmt.Println("There was a problem receiving messege,", err)
			continue
		}
		fmt.Println(strings.Trim(string(rbuff), "\x00"))
	}
}

func establishSender(conn net.TCPConn, c chan int) {
	scn := bufio.NewScanner(os.Stdin)
	for {
		scn.Scan()
		text := scn.Text()
		_, err := conn.Write([]byte(text)); if err != nil {
			fmt.Println("There was a problem sending the messege,", err)
			continue
		}
	}
}