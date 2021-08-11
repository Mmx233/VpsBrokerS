package dao

import (
	"gorm.io/gorm"
)

type Event struct {
	ID      uint
	Type    string
	Time    int64  `gorm:"index"`
	Ip      string `gorm:"index,not null"`
	DownNum uint
	AllNum  uint
	Vps     Vps `gorm:"foreignKey:Ip,references:Ip,constraint:OnUpdate:CASCADE"`
}

func (a *Event) Insert(tx *gorm.DB) error {
	return defaultInsert(tx, a)
}

func (a Event) CountDown() (int64, error) {
	return defaultCounter(&Event{Type: "down", Ip: a.Ip})
}

func (a Event) CountUp() (int64, error) {
	return defaultCounter(&Event{Type: "up", Ip: a.Ip})
}
