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
	buffer := make([]byte, 1024)
	for {
		total, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:total]))
	}
}
