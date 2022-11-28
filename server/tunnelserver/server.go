package tunnelserver

import (
	"fmt"
	"lunelerG/service/session"
	"net"

	log "github.com/sirupsen/logrus"
)

type TunnelServer struct {
	Addr        *net.TCPAddr
	Distributor *session.Distributor
}

func (t *TunnelServer) Start() {
	listener, err := net.ListenTCP("tcp", t.Addr)
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
		log.Printf("New connection to tunnel server from :: %s \n", conn.RemoteAddr().String())
		t.Distributor.AddTunnelSession(session.NewTunnelSession(conn, make(chan session.Message)))
	}
}

func NewTunnelServer(port int, distributer session.Distributor) *TunnelServer {
	return &TunnelServer{
		Addr: &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: port,
		},
		Distributor: &distributer,
	}
}
