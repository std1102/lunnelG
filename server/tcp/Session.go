package tcp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"lunelerG/server/service"
	"lunelerG/utils"
	"net"
)

type Session struct {
	conn    *net.TCPConn
	id      string
	channel chan service.Message
}

func (sess *Session) GetChannel() chan service.Message {
	return sess.channel
}

func (sess *Session) HandleRecvData(recv chan service.Message) {
	//TODO implement me
	panic("implement me")
}

func (sess *Session) HandleData(data []byte) {
	//TODO implement me
	panic("implement me")
}

func NewSession(conn *net.TCPConn) Session {
	var session Session
	session.conn = conn
	session.channel = make(chan service.Message)
	session.GenerateId()
	return session
}

func (sess *Session) GetIp() string {
	return sess.conn.RemoteAddr().String()
}

func (sess *Session) GetId() string {
	return sess.id
}

func (sess *Session) GenerateId() {
	var hexId utils.HexId
	hexId.RandomString(service.ID_LENGTH)
	ipId := sess.conn.RemoteAddr().String() + " " + hexId.Id
	sess.id = ipId
}

func (sess *Session) Start() {
	defer sess.Close()
	buffer := make([]byte, 1024)
	for {
		total, err := sess.conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Errorln("Cannot read data \n")
			break
		}
		fmt.Println(string(buffer[:total]))
	}
}

func (sess *Session) Close() {
	err := sess.conn.Close()
	if err != nil {
		log.Error("Error when close connection")
		fmt.Println(err)
	}
}
