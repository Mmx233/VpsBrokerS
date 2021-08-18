package modules

import (
	"github.com/Mmx233/VpsBrokerS/util"
	"github.com/gorilla/websocket"
	"sync"
)

type pool struct {
	Clients *sync.Map
}

type clientInfo struct {
	Lock    sync.RWMutex
	Conn    *websocket.Conn
	Name    string
	Port    uint
	DownNum uint //客户端连接中有几台发生down
}

// Pool  客户端连接池
var Pool = pool{
	Clients: &sync.Map{},
}

// Add 加入连接
func (a *pool) Add(ip string, conn *websocket.Conn, name string, port uint) bool {
	d, ok := a.Clients.LoadOrStore(ip, &clientInfo{
		Conn: conn,
		Name: name,
		Port: port,
	})
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		t.Conn = conn
		t.Name = name
		t.Port = port
		t.Lock.Unlock()
	} else {
		a.SendListInfoAll()
	}

	util.Event.NewClientConn(name)

	return !ok
}

// Lose 连接失效
func (a *pool) Lose(ip string) {
	d, ok := a.Clients.Load(ip)
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		_ = t.Conn.Close()
		t.Conn = nil
		util.Event.LostClient(t.Name)
		t.Lock.Unlock()
	}

	a.SendListInfoAll()
}

// ClientUp 上线记录
func (a *pool) ClientUp(ip string) uint {
	d, _ := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.DownNum--
	return t.DownNum
}

// ClientDown 掉线记录
func (a *pool) ClientDown(ip string) uint {
	d, _ := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.DownNum++
	return t.DownNum
}

func (a *pool) RecoverDownNum(ip string, num int64) {
	d, ok := a.Clients.Load(ip)
	if ok {
		t := d.(*clientInfo)
		t.Lock.Lock()
		t.DownNum = uint(num)
		t.Lock.Unlock()
	}
}

// Load 读取连接
func (a *pool) Load(ip string) (*websocket.Conn, uint, bool) {
	d, ok := a.Clients.Load(ip)
	t := d.(*clientInfo)
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	return t.Conn, t.Port, ok
}

// GetListInfo 获取客户端列表
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

// SendListInfoAll 向所有客户端发送列表
func (a *pool) SendListInfoAll() {
	a.Clients.Range(func(ip, t interface{}) bool {
		d := t.(*clientInfo)
		i := 0
		for {
			if d.Conn == nil {
				continue
			}
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

// Len 池长度
func (a *pool) Len() (i int) {
	a.Clients.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return
}
