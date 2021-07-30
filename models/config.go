package models

type Config struct {
	Settings Settings
	Mysql    Mysql
}

type Settings struct {
	Port uint
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     uint
	Database string
	Arg      string
}
