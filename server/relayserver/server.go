package relayserver

import (
	"fmt"
	"lunelerG/service/session"
	"net"

	log "github.com/sirupsen/logrus"
)

type RelayServer struct {
	Addr        *net.TCPAddr
	Distributor *session.Distributor
}

func (r *RelayServer) Start() {
	listener, err := net.ListenTCP("tcp", r.Addr)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Cannot bind server")
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			log.Errorln("Error in making connection")
			continue
		}
		tunnelSessionId := r.Distributor.GetSessionTunnelId()
		requestSession := session.NewRequestSession(
			*conn,
			r.Distributor.HeaderChannel,
			r.Distributor.TunnelSessions[tunnelSessionId].Transmitter,
			tunnelSessionId,
		)
		go requestSession.Start()
	}
}

func NewRelaylServer(port int, distributer session.Distributor) *RelayServer {
	return &RelayServer{
		Addr: &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: port,
		},
		Distributor: &distributer,
	}
}
