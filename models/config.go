package models

type Config struct {
	Settings Settings
	Mysql    Mysql
	Redis    Redis
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

type Redis struct {
	Addr     string
	Password string
	DB       int
}
