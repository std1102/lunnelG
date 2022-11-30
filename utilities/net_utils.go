package util

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Errorln("CANNOT CLOSE CONNECTION " + err.Error())
			return
		}
	}(conn)
	reader := bufio.NewReader(conn)
	netData := ""
	for {
		message, err := reader.ReadString('\r')
		if err != nil {
			break
		}
		netData += message
	}
	fmt.Print(netData)
	conn.Write([]byte("CON CHO"))
}
