package controllers

import (
	"github.com/gorilla/websocket"
	"sync"
)

type pool struct {
	Clients *sync.Map
}

type clientInfo struct {
	Lock sync.RWMutex
	Conn *websocket.Conn
	Port uint
}

var Pool = pool{
	Clients: &sync.Map{},
}

func (a *pool) Add(ip string, conn *websocket.Conn, port uint) {
	d, ok := a.Clients.LoadOrStore(ip, &clientInfo{
		Conn: conn,
		Port: port,
	})
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		t.Conn = conn
		t.Port = port
		t.Lock.Unlock()

	}
}

func (a *pool) Down(ip string) {
	d, ok := a.Clients.Load(ip)
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		t.Conn = nil
		t.Lock.Unlock()
	}
}

func (a *pool) Load(ip string) (*websocket.Conn, uint, bool) {
	d, ok := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.Conn, t.Port, ok
}

func (a *pool) SendListInfo() {

}
