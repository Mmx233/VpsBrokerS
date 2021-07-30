package dao

type Nickname struct {
	Ip   string `gorm:"unique,index"`
	Name string
}
