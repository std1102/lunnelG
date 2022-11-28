package session

import (
	log "github.com/sirupsen/logrus"
)

type Distributor struct {
	TunnelSessions map[string]TunnelSession
	HeaderChannel  chan Message
}

func (d *Distributor) GetSessionTunnelId() string {
	min := 9999999999999999
	var tunnelSessionMin TunnelSession
	for _, tunnelSession := range d.TunnelSessions {
		if len(tunnelSession.Transmitter) < min {
			tunnelSessionMin = tunnelSession
			min = len(tunnelSession.Transmitter)
		}
	}
	return tunnelSessionMin.Id
}

func (d *Distributor) Distribute() {
	for {
		revcMessage := <-d.HeaderChannel
		tunnelSession, isPresent := d.TunnelSessions[revcMessage.TunnelsessionId]
		if !isPresent {
			log.Errorln("Cannot find tunnel session ID, it maybe deleted")
			continue
		}
		tunnelSession.Transmitter <- revcMessage
	}
}

func NewDistributor() *Distributor {
	return &Distributor{
		TunnelSessions: make(map[string]TunnelSession),
		HeaderChannel:  make(chan Message),
	}
}

func (d *Distributor) GetTunnelSession(id string) TunnelSession {
	return d.TunnelSessions[id]
}

func (d *Distributor) AddTunnelSession(session *TunnelSession) {
	d.TunnelSessions[session.Id] = *session
	go session.Start()
}

func (d *Distributor) RemoveChannel(id string) {
	//TODO implement me
	panic("implement me")
}
