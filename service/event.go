package service

import (
	"github.com/Mmx233/VpsBrokerS/models/dao"
	"time"
)

type event struct{}

var Event event

func (*event) Up(ip string, Time int64) error {
	t := dao.Event{
		Type: "up",
		Time: Time,
		Ip:   ip,
	}
	return t.Insert(nil)
}

func (*event) Down(ip string) error {
	t := dao.Event{
		Type: "down",
		Time: time.Now().UnixNano(),
		Ip:   ip,
	}
	return t.Insert(nil)
}
