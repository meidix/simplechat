package chat

import (
	"fmt"
	"meidix/chatserv/connection"
	"net"
	"os"
)

func MessegeListener(u string, conn net.TCPConn, ch chan Messege) {
	for {
		text, err := connection.RecvTxtMsg(&conn); if err != nil {
			fmt.Println("There was a problem receiving messege")
			os.Exit(1)
		}
		ch <- Messege{u, text}
	}
}


func MessegeDispatcher(u map[string]net.TCPConn, ch chan Messege) {
	for m := range ch {
		for key, conn := range u {
			if key != m.sender {
				err := connection.SendTxtMsg(&conn, m.toString()); if err != nil {
					fmt.Println("there was a problem sending messeges")
					delete(u, key)
				}
			}
		}
	}
}

func SetupChat(key string, conn *net.TCPConn, u map[string]net.TCPConn, ch chan Messege) {
	u[key] = *conn
	go MessegeListener(key, *conn, ch)
	err := connection.SendTxtMsg(conn, "Welcome to This Chat")
	if err  != nil {
		fmt.Println("There was a problem in sending a messege", err)
	}
}
