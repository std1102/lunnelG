package tcp

import (
	"fmt"
	"log"
	"net"
)

type Session struct {
	conn *net.TCPConn
}

func NewSession(conn *net.TCPConn) *Session {
	return &Session{conn: conn}
}

func (sess Session) Start() {
	log.Printf("CONNECTION FROM :: %s", sess.conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	for {
		total, err := sess.conn.Read(buffer)
		if err != nil {
			log.Fatal("CANNOT READ DATA \n")
			break
		}
		fmt.Println(string(buffer[:total]))
	}
}
