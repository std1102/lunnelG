package main

import (
	"lunelerG/client"
	"lunelerG/server/relayserver"
	"lunelerG/server/tunnelserver"
	"lunelerG/service/session"
	"time"
)

func main() {
	distributor := session.NewDistributor()
	go distributor.Distribute()
	rServer := relayserver.NewRelaylServer(9999, *distributor)
	go rServer.Start()
	tServer := tunnelserver.NewTunnelServer(7777, *distributor)
	go tServer.Start()
	gClient := client.Client{
		IPAddr: "127.0.0.1:7777",
	}
	go gClient.Start()
	time.Sleep(10000000 * time.Second)
}
