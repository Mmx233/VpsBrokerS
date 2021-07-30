package dao

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerS/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

//连接mysql数据库
func connect(username string, password string, host string, port uint, database string, arg string) *gorm.DB {
	var dbConfig gorm.Config

	db, e := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		username,
		password,
		host,
		port,
		database,
		arg,
	)), &dbConfig)
	if e != nil {
		log.Fatalln("Mysql数据库初始化失败\n", e)
	}
	return db
}

func init() {
	//数据库初始化
	db = connect(
		global.Config.Mysql.Username,
		global.Config.Mysql.Password,
		global.Config.Mysql.Host,
		global.Config.Mysql.Port,
		global.Config.Mysql.Database,
		global.Config.Mysql.Arg,
	)
	//连接池设置
	if sqlDB, err := db.DB(); err != nil {
		log.Fatalln("sqlDB获取失败\n", err)
	} else {
		sqlDB.SetConnMaxLifetime(time.Hour * 5)
	}
	//自动迁移
	if err := db.AutoMigrate(); err != nil {
		log.Fatalln(err)
	}
}
