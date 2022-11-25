package distributer

import "net"

type Message struct {
	Id   string
	Data net.Conn
}

func NewMessage(id string, data net.Conn) *Message {
	return &Message{Id: id, Data: data}
}
