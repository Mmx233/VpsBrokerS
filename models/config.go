package models

type Config struct {
	Mysql Mysql
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     uint
	Database string
	Arg      string
}
