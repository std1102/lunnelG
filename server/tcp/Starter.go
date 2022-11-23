package tcp

import (
	"fmt"
	"log"
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

// Bind handle ssl here (QUICC)
func (tcpServer TcpServer) Bind() error {
	listener, err := net.ListenTCP("tcp", tcpServer.Addr)
	if err != nil {
		log.Fatalf("Cannot bind address")
		fmt.Println(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		defer conn.Close()
		session := NewSession(conn)
		go session.Start()
	}
}
