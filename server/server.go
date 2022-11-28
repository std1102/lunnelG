package server

import "net"

type Server struct {
	Addr *net.TCPAddr
}

func NewServer(port int) *Server {
	return &Server{
		Addr: &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: port,
		},
	}
}

type ServerAction interface {
	Start()
}
