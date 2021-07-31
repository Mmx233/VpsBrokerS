package dao

import "gorm.io/gorm"

type Vps struct {
	ID     uint
	Ip     string `gorm:"unique,index"`
	Name   string
	Online bool `gorm:"index"`
}

func (a Vps) Get() ([]Vps, error) {
	var t = make([]Vps, 0)
	return t, defaultGetWhitQuery(&a, &t)
}

func (a *Vps) Insert(tx *gorm.DB) error {
	return defaultInsert(tx, a)
}

func (a Vps) Exist() bool {
	return defaultExist(&a)
}

func (a *Vps) Find() error {
	return defaultFind(a)
}

func (a *Vps) UpdateName() error {
	return db.Where(Vps{Ip: a.Ip}).Updates(a).Error
}
