package modules

import (
	"sync"
	"time"
)

type onlineStack struct {
	Lock     sync.RWMutex
	Timeline time.Time
}

func (a *onlineStack) Continue() {
	a.Lock.Lock()
	a.Timeline = time.Now()
	a.Lock.Unlock()
}

func (a *onlineStack) Seconds() float64 {
	a.Lock.RLock()
	t := time.Since(a.Timeline).Seconds()
	a.Lock.RUnlock()
	return t
}

type online struct {
	data sync.Map
}

var Online online

func (a *online) Up(ip string) {
	a.data.Store(ip, &onlineStack{Timeline: time.Now()})
}

func (a *online) Down(ip string) {
	d, ok := a.data.LoadAndDelete(ip)
	if ok {
		t := d.(*onlineStack)
		t = nil
		_ = t
	}
}

func (a *online) Continue(ip string) {
	d, ok := a.data.Load(ip)
	if !ok {
		a.Up(ip)
		return
	}
	d.(*onlineStack).Continue()
}
