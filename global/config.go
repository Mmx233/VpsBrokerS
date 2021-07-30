package global

import (
	"github.com/Mmx233/VpsBrokerS/models"
	"github.com/Mmx233/config"
	"log"
	"os"
)

var Config models.Config

func init() {
	if e := config.Load(config.Options{
		Config: &Config,
		Default: &models.Config{
			Mysql: models.Mysql{
				Host: "127.0.0.1",
				Port: 3306,
				Arg:  "charset=utf8mb4&parseTime=true&loc=Local",
			},
		},
		FillDefault: true,
	}); e != nil {
		if config.IsNew(e) {
			log.Println("已生成配置文件'Config.json'，请编辑后再次启动")
			os.Exit(1)
		}

		log.Fatalln("配置文件初始化失败：\n", e)
	}
}
