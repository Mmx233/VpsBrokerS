package dao

import (
	"github.com/Mmx233/daoUtil"
	"gorm.io/gorm"
)

type Vps struct {
	ID uint `json:"id"`
	gorm.DeletedAt
	Ip   string `json:"ip" gorm:"unique,index"`
	Name string `json:"name"`
	Port uint   `json:"port"`
}

func (a Vps) Get() ([]Vps, error) {
	var t = make([]Vps, 0)
	return t, daoUtil.DefaultGetWhitQuery(&a, &t)
}

func (a *Vps) Insert() error {
	return daoUtil.DefaultInsert(a)
}

func (a Vps) Exist() bool {
	return daoUtil.DefaultExist(&a)
}

func (a *Vps) Find() error {
	return daoUtil.DefaultFind(a)
}

func (a *Vps) Update() error {
	return db.Where(Vps{Ip: a.Ip}).Updates(a).Error
}

func (a Vps) Delete() error {
	return daoUtil.DefaultDelete(&a)
}
