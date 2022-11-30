package session

import (
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
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
	result := make([]byte, 0)
	buffer := make([]byte, 256)
	result = append(result)
	byteId, _ := hex.DecodeString(requestId)
	for {
		total, err := newCommingConn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		result = append(result, buffer[:total]...)
		if total < 256 {
			break
		}
	}
	fmt.Println(len(result))
	log.Printf("Sending data to :: %s \n", responseConn.RemoteAddr())
	for i := 0; i < int(math.Ceil(float64(len(result)/8))); i++ {
		sendBuffer := append(byteId, result[i:i+6]...)
		responseConn.Write(sendBuffer)
	}
	responseConn.Write(append(byteId, make([]byte, 0, 6)...))
}
