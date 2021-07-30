package dao

type Event struct {
	ID       uint
	Time     int64    `gorm:"index"`
	Ip       string   `gorm:"index,not null"`
	Nickname Nickname `gorm:"foreignKey:Ip,references:Ip,constraint:OnUpdate:CASCADE"`
}
