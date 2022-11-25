package session

import (
	"io"
	"lunelerG/service/distributer"
	hexid "lunelerG/utilities"
	"net"
	"regexp"
	"strings"
)

const RESPONSE = "RESPONSE"
const REQUEST = "REQUEST"

type RequestSession struct {
	id         string
	conn       net.TCPConn
	trasmitter chan distributer.Message
}

func (r RequestSession) Start() {
	message := distributer.NewMessage(r.id, &r.conn)
	r.trasmitter <- *message
	// wait and response
	for {
		recvMessage := <-r.trasmitter
		if filterResponse(r.id, recvMessage.Id) {
			io.Copy(&r.conn, recvMessage.Data)
			return
		}
	}
}

func NewRequestSession(conn net.TCPConn, transmitter chan distributer.Message) *RequestSession {
	return &RequestSession{
		id:         hexid.New(hexid.DEFAULT_LENGTH),
		conn:       conn,
		trasmitter: transmitter,
	}
}

func filterResponse(id, idResponse string) bool {
	if m, _ := regexp.MatchString("^RESPONSE.+", idResponse); !m {
		return false
	}
	stringArr := strings.Split(idResponse, " ")
	if stringArr[0] == RESPONSE && stringArr[1] == id {
		return true
	}
	return false
}
