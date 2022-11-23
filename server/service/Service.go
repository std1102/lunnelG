package service

import "net"

type Session interface {
	Start(conn net.Conn) error
	Close() error
}
