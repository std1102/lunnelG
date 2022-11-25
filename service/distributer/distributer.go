package distributer

import "lunelerG/service/session"

type Distributor struct {
	TunnelSession map[string]session.TunnelSession
}

func (d Distributor) GetTunnelSession(id string) session.TunnelSession {
	return d.TunnelSession[id]
}

func (d Distributor) AddTunnelSession(session session.TunnelSession) {
	d.TunnelSession[session.Id] = session
}

func (d Distributor) RemoveChannel(id string) {
	//TODO implement me
	panic("implement me")
}

func (d Distributor) AttachRequestSession(idTunnelSession string, request session.RequestSession) {

}

func NewDistributor() *Distributor {
	return &Distributor{
		TunnelSession: make(map[string]session.TunnelSession),
	}
}

type DistributorAction interface {
	AddTunnelSession(session session.TunnelSession)
	RemoveChannel(id string)
	AttachRequestSession(idTunnelSession string, request session.RequestSession)
	GetTunnelSession(id string) session.TunnelSession
}
