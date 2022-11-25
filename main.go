package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:6789")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func forward(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

func handleConnection(c net.Conn) {

	log.Println("Connection from : ", c.RemoteAddr())

	remote, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	// go routines to initiate bi-directional communication for local server with a
	// remote server
	go forward(c, remote)
	go forward(remote, c)
}
