package connection

import (
	"net"
	"strings"
)

func SendTxtMsg(conn *net.TCPConn, msg string) error {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func RecvTxtMsg(conn *net.TCPConn) (string, error) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		return "", err
	} else {
		text := strings.Trim(string(buff), "\x00")
		return text, nil
	}
}