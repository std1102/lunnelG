package client

import (
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	util "lunelerG/utilities"
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
	index := 0
	for {
		index++
		buffer := make([]byte, 8)
		_, err := caller.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		requestId := hex.EncodeToString(buffer[0:2])
		if len(r.RequestSet[requestId]) == 0 {
			r.RequestSet[requestId] = buffer[2:8]
		} else {
			tempArr := append(r.RequestSet[requestId], buffer[2:8]...)
			r.RequestSet[requestId] = tempArr
			if compareSlice(buffer[2:8], make([]byte, 6)) {
				request, dErr := net.Dial("tcp", "127.0.0.1:8081")
				if dErr != nil {
					fmt.Println("ERROR WHEN CONNECT TARGER " + dErr.Error())
				}
				requestArr := backTrimSlice(tempArr)
				request.Write(requestArr)
				util.HandleConnection(request)
			}
		}
	}
}

func compareSlice(slice1 []byte, slice2 []byte) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	result := true
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			result = false
		} else {
			result = true
		}
	}
	return result
}

func backTrimSlice(data []byte) []byte {
	sliceIndex := 0
	for i := len(data) - 1; i > 0; i-- {
		if (data)[i] == 0 {
			sliceIndex++
		}
		break
	}
	data = data[0 : len(data)-sliceIndex]
	return data
}
