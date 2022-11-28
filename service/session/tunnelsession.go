package session

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

type TunnelSession struct {
	Id          string
	Conn        *net.Conn
	Transmitter chan Message
}

func (t TunnelSession) Start() {
	for {
		recvMessage := <-t.Transmitter
		go handleRequest(recvMessage.Data, *t.Conn, t.Transmitter)
	}
}

func NewTunnelSession(conn net.Conn, trasmitter chan Message) *TunnelSession {
	return &TunnelSession{Conn: &conn, Transmitter: trasmitter}
}

func handleRequest(newCommingConn net.Conn, responseConn net.Conn, transmiter chan Message) {
	defer newCommingConn.Close()
	result := make([]byte, 0)
	for {
		buffer := make([]byte, 1024)
		total, err := newCommingConn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		result = append(result, buffer[:total]...)
	}
	log.Printf("Sending data to :: %s \n", responseConn.RemoteAddr())
	responseConn.Write(result)
}
