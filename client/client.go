package client

import (
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"reflect"
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
	IPAddr     string
	RequestSet map[string][]byte
}

func NewClient(IPAddr string) *Client {
	return &Client{IPAddr: IPAddr, RequestSet: make(map[string][]byte)}
}

func (r *Client) Start() {
	caller, err := net.Dial("tcp", r.IPAddr)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot bind server")
	}
	defer caller.Close()
	request, dErr := net.Dial("tcp", "127.0.0.1:8081")
	if dErr != nil {
		fmt.Println("ERROR WHEN CONNECT TARGER " + dErr.Error())
	}
	for {
		buffer := make([]byte, 8)
		_, err := caller.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		requestId := hex.EncodeToString(buffer[0:2])
		if len(r.RequestSet[requestId]) == 0 {
			r.RequestSet[requestId] = buffer[2:5]
		} else {
			tempArr := append(r.RequestSet[requestId], buffer[2:5]...)
			r.RequestSet[requestId] = tempArr
			if reflect.DeepEqual(buffer[2:5], make([]byte, 0, 6)) {
				request.Write(tempArr)
			}
		}
	}
}
