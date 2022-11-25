package server

import "net"

type Server struct {
	Addr *net.TCPAddr
}

type ServerAction interface {
}
