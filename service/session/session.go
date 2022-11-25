package session

import (
	"io"
	"net"
)

type Session interface {
	Start()
}

func Forward(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
