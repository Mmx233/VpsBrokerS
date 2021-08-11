package service

import (
	"github.com/Mmx233/VpsBrokerS/models/dao"
	"github.com/Mmx233/VpsBrokerS/service/modules"
)

type event struct{}

// Event 事件记录器
var Event event

func (*event) log(Type string, ip string, Time int64, downNum uint) error {
	t := dao.Event{
		Type:    Type,
		Time:    Time,
		Ip:      ip,
		DownNum: downNum,
		AllNum:  uint(modules.Pool.Len()),
	}
	return t.Insert(nil)
}

// Up 连接恢复事件
func (a *event) Up(ip string, Time int64) error {
	return a.log("up", ip, Time, modules.Pool.ClientUp(ip))
}

// Down 掉线事件
func (a *event) Down(ip string, Time int64) error {
	return a.log("down", ip, Time, modules.Pool.ClientDown(ip))
}
