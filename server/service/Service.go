package service

import (
	"log"
	"reflect"
)

const ID_LENGTH = 10

type Session interface {
	Start()
	Close()
	GetIp() string
	GetId() string
	GenerateId()
	GetChannel() chan Message
	HandleRecvData(recv chan Message)
	// More function to implement SSL
}

func NewSessionStore() SessionStore {
	return SessionStore{
		storage: make(map[string]Session),
	}
}

type SessionMap interface {
	Add(session Session, recv chan Message)
	Remove(id string)
	ContainsKey(key string) bool
	GetAllKey() []string
	GetSession(key string) Session
}

func containsKey(m, k interface{}) bool {
	v := reflect.ValueOf(m).MapIndex(reflect.ValueOf(k))
	return v != reflect.Value{}
}

type SessionStore struct {
	storage map[string]Session
}

func (sessionStore *SessionStore) GetAllKey() []string {
	keys := make([]string, len(sessionStore.storage))
	for k := range sessionStore.storage {
		keys = append(keys, k)
	}
	return keys
}

func (sessionStore *SessionStore) GetSession(key string) Session {
	if !sessionStore.ContainsKey(key) {
		return nil
	}
	return sessionStore.storage[key]
}

func (sessionStore *SessionStore) ContainsKey(key string) bool {
	return containsKey(sessionStore.storage, key)
}

func (sessionStore *SessionStore) Add(session Session) {
	sessionStore.storage[session.GetId()] = session
	log.Printf("Added connection from %s session map, it will start soon", session.GetId())
	go session.Start()
}

func (sessionStore *SessionStore) Remove(id string) {
	if sessionStore.ContainsKey(id) {
		return
	}
	delete(sessionStore.storage, id)
}
