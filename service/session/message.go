package session

import "net"

type Message struct {
	RequestId       string
	TunnelsessionId string
	Data            net.Conn
}

func NewMessage(requestId string, sessionId string, data net.Conn) *Message {
	return &Message{RequestId: requestId, TunnelsessionId: sessionId, Data: data}
}
