package service

import (
	"regexp"
	"strings"
	"sync"
)

const QUIT = "QUIT"

type Message struct {
	sessionId string
	data      []byte
}

type DataDistrubute interface {
	Add(session Session)
	RemoveChannel(sessionId string)
	SendData(message Message)
	Distribute()
}

type Distributer struct {
	mutex        sync.Mutex
	headChannel  chan Message
	SessionStore SessionStore
}

// Impl of DataDistrubute
func (d *Distributer) RemoveChannel(sessionId string) {
	if !d.SessionStore.ContainsKey(sessionId) {
		return
	}
	d.SessionStore.Remove(sessionId)
}

func NewDistributer() *Distributer {
	return &Distributer{
		headChannel:  make(chan Message),
		SessionStore: NewSessionStore(),
	}
}

// Impl of DataDistrubute
// Add a channel to session
func (d *Distributer) AddChannel(sessionId string) {
	if !d.SessionStore.ContainsKey(sessionId) {
		return
	}
}

// Impl of DataDistrubute
func (d *Distributer) SendData(message Message) {
	if !d.SessionStore.ContainsKey(message.sessionId) {
		return
	}
	d.mutex.Lock()
	d.SessionStore.storage[message.sessionId].GetChannel() <- message
	d.mutex.Unlock()
}

// Impl of DataDistrubute
func (d *Distributer) Distribute() {
	for {
		message := <-d.headChannel
		if ismatch, _ := regexp.MatchString("^QUIT.*", message.sessionId); ismatch {
			sessionId := strings.Split(message.sessionId, " ")[1]
			d.RemoveChannel(sessionId)
			continue
		}
		d.mutex.Lock()
		d.SendData(message)
		d.mutex.Unlock()
	}
}
