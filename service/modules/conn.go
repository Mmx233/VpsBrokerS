package modules

import (
	"github.com/gorilla/websocket"
	"sync"
)

type pool struct {
	Clients *sync.Map
}

type clientInfo struct {
	Lock    sync.RWMutex
	Conn    *websocket.Conn
	Port    uint
	DownNum uint
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

	a.SendListInfoAll()
}

func (a *pool) Lose(ip string) {
	d, ok := a.Clients.Load(ip)
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		t.Conn = nil
		t.Lock.Unlock()
	}

	a.SendListInfoAll()
}

func (a *pool) ClientUp(ip string) uint {
	d, _ := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.DownNum--
	return t.DownNum
}

func (a *pool) ClientDown(ip string) uint {
	d, _ := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.DownNum++
	return t.DownNum
}

func (a *pool) Load(ip string) (*websocket.Conn, uint, bool) {
	d, ok := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.Conn, t.Port, ok
}

func (a *pool) GetListInfo() map[string]uint {
	var data = make(map[string]uint)
	a.Clients.Range(func(ip, t interface{}) bool {
		d := t.(*clientInfo)
		d.Lock.RLock()
		data[ip.(string)] = d.Port
		d.Lock.RUnlock()
		return true
	})

	return data
}

func (a *pool) SendListInfoAll() {
	a.Clients.Range(func(ip, t interface{}) bool {
		d := t.(*clientInfo)
		i := 0
		for {
			if d.Conn.WriteJSON(a.GetListInfo()) == nil {
				break
			} else {
				if i >= 5 {
					break
				}
				i++
			}
		}

		return true
	})
}

func (a *pool) Len() (i int) {
	a.Clients.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return
}
