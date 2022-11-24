package tcp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lunelerG/server/service"
	"net"
	"time"
)

const (

	// Time to wait before retrying a failed Accept().
	acceptRetryWait = 100 * time.Millisecond
)

type TcpServer struct {
	Addr *net.TCPAddr
}

func (tcpServer *TcpServer) Bind(port int) {
	tcpServer = initServer(port)
	listener, err := net.ListenTCP("tcp", tcpServer.Addr)
	if err != nil {
		log.Fatalf("Cannot bind port")
		fmt.Println(err)
	}
	defer listener.Close()

	// Create distributer
	distributer := service.NewDistributer()
	go distributer.Distribute()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		session := NewSession(conn)
		distributer.SessionStore.Add(&session)
	}
}

// init server object
func initServer(port int) *TcpServer {
	address := net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	}
	return &TcpServer{
		Addr: &address,
	}
}
