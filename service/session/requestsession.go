package session

import (
	"io"
	"log"
	hexid "lunelerG/utilities"
	"net"
	"regexp"
	"strings"
)

const RESPONSE = "RESPONSE"
const REQUEST = "REQUEST"

type RequestSession struct {
	id                   string
	tunnelSessionId      string
	conn                 *net.TCPConn
	headChannel          chan Message
	tunnelSessionChannel chan Message
}

func (r RequestSession) Start() {
	log.Printf("New comming request to relay server from :: %s \n", r.conn.RemoteAddr().String())
	message := NewMessage(r.id, r.tunnelSessionId, r.conn)
	r.headChannel <- *message
	// wait and response
	for {
		recvMessage := <-r.tunnelSessionChannel
		if filterResponse(r.id, recvMessage.RequestId) {
			defer r.conn.Close()
			io.Copy(r.conn, recvMessage.Data)
			return
		}
	}
}

func NewRequestSession(conn net.TCPConn, headChannel chan Message, tunnelSessionChannel chan Message, tunnelSessionId string) *RequestSession {
	return &RequestSession{
		id:                   hexid.New(hexid.DEFAULT_LENGTH),
		tunnelSessionId:      tunnelSessionId,
		conn:                 &conn,
		headChannel:          headChannel,
		tunnelSessionChannel: tunnelSessionChannel,
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
