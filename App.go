package main

import "lunelerG/server/tcp"

func main() {
	var server tcp.TcpServer
	server.Bind(7777)
}
