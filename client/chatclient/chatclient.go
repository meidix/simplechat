package chatclient

import (
	"bufio"
	"fmt"
	"meidix/chatclient/connection"
	"net"
	"os"
)

func EstablishReceiver(conn net.TCPConn, c chan int) {
	for {
		msg, err := connection.RecvTxtMsg(&conn)
		if err != nil {
			fmt.Println("There was a problem receiving messege,", err)
			os.Exit(1)
		}
		fmt.Println(msg)
	}
}

func EstablishSender(conn net.TCPConn, c chan int) {
	scn := bufio.NewScanner(os.Stdin)
	for {
		scn.Scan()
		text := scn.Text()
		err := connection.SendTxtMsg(&conn, text)
		if err != nil {
			fmt.Println("There was a problem sending the messege,", err)
			os.Exit(1)
		}
	}
}