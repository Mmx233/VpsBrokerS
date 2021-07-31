package dao

type Vps struct {
	Ip     string `gorm:"unique,index"`
	Name   string
	Online bool
}
