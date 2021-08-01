package dao

import "gorm.io/gorm"

type Event struct {
	ID   uint
	Type string
	Time int64  `gorm:"index"`
	Ip   string `gorm:"index,not null"`
	Vps  Vps    `gorm:"foreignKey:Ip,references:Ip,constraint:OnUpdate:CASCADE"`
}

func (a *Event) Insert(tx *gorm.DB) error {
	return defaultInsert(tx, a)
}
