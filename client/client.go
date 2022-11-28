package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

func forward(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

func handleConnection(addr string, c net.Conn) {

	log.Println("Connection from : ", c.RemoteAddr())

	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// go routines to initiate bi-directional communication for local server with a
	// remote server
	go forward(c, remote)
	go forward(remote, c)
}

type Client struct {
	IPAddr string
}

func (r *Client) Start() {
	caller, err := net.Dial("tcp", r.IPAddr)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot bind server")
	}
	//defer caller.Close()
	for {
		buffer := make([]byte, 1024)
		total, err := caller.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer[:total]))
	}
}
