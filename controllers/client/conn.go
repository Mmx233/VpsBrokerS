package controllers

import (
	"github.com/gorilla/websocket"
	"sync"
)

type pool struct {
	Conn *sync.Map
}

var Pool = pool{
	Conn: &sync.Map{},
}

func (a *pool) Add(ip string, conn *websocket.Conn) {
	a.Conn.Store(ip, conn)
}

func (a *pool) Rm(ip string) {
	t, ok := a.Conn.LoadAndDelete(ip)
	if ok {
		t = nil
		_ = t
	}
}

func (a *pool) Load(ip string) (*websocket.Conn, bool) {
	conn, ok := a.Conn.Load(ip)
	return conn.(*websocket.Conn), ok
}
