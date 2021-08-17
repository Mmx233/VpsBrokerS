package dao

import "gorm.io/gorm"

type Vps struct {
	ID   uint   `json:"id"`
	Ip   string `json:"ip" gorm:"unique,index"`
	Name string `json:"name"`
	Port uint   `json:"port"`
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

func (a *Vps) Update() error {
	return db.Where(Vps{Ip: a.Ip}).Updates(a).Error
}
