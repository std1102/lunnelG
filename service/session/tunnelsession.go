package session

import (
	"lunelerG/service/distributer"
	"net"
)

type TunnelSession struct {
	Id         string
	Conn       net.Conn
	Trasmitter chan distributer.Message
}

func (t TunnelSession) Start() {
	for {
		recvMessage := <-t.Trasmitter
		Forward(recvMessage.Data, t.Conn)
		sendMessage := distributer.NewMessage(RESPONSE+" "+recvMessage.Id, nil)
		Forward(t.Conn, sendMessage.Data)
		t.Trasmitter <- *sendMessage
	}
}

func NewTunnelSession(conn net.Conn, trasmitter chan distributer.Message) *TunnelSession {
	return &TunnelSession{Conn: conn, Trasmitter: trasmitter}
}
