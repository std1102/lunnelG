package main

import (
	"lunelerG/server/tcp"
	"net"
)

func main() {

	address := net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 7777,
	}

	newConn := tcp.TcpServer{
		Addr: &address,
	}

	newConn.Bind()
}
