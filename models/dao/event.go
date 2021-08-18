package dao

import (
	"github.com/Mmx233/daoUtil"
)

type Event struct {
	ID      uint
	Type    string
	Time    int64  `gorm:"index"`
	Ip      string `gorm:"index,not null"`
	DownNum uint
	AllNum  uint
}

func (a *Event) Insert() error {
	return daoUtil.DefaultInsert(a)
}

func (a Event) CountDown() (int64, error) {
	return daoUtil.DefaultCounter(&Event{Type: "down", Ip: a.Ip})
}

func (a Event) CountUp() (int64, error) {
	return daoUtil.DefaultCounter(&Event{Type: "up", Ip: a.Ip})
}
