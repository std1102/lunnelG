package session

import (
	"encoding/hex"
	"fmt"
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
		go handleRequest(recvMessage.RequestId, recvMessage.Data, *t.Conn, t.Transmitter)
	}
}

func NewTunnelSession(conn net.Conn, trasmitter chan Message) *TunnelSession {
	return &TunnelSession{Conn: &conn, Transmitter: trasmitter}
}

func handleRequest(requestId string, newCommingConn net.Conn, responseConn net.Conn, transmiter chan Message) {
	buffer := make([]byte, 6)
	byteId, _ := hex.DecodeString(requestId)
	index := 0
	for {
		index++
		total, err := newCommingConn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				responseConn.Write(append(byteId, make([]byte, 0, 6)...))
				break
			}
			fmt.Println(err)
		}
		sendData := append(byteId, buffer[:total]...)
		responseConn.Write(sendData)
		if total < 6 {
			log.Printf("Sent all data to :: %s", responseConn.RemoteAddr())
			responseConn.Write(append(byteId, make([]byte, 0, 6)...))
			break
		}
	}
}
