package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/VpsBrokerS/controllers/modules"
	"github.com/Mmx233/VpsBrokerS/models/form"
	"github.com/Mmx233/VpsBrokerS/service"
	"github.com/gin-gonic/gin"
)

func Init(c *gin.Context) {
	var f form.VpsInit
	if e := c.ShouldBind(&f); e != nil {
		controllers.CallBack.Error(c, 2)
		return
	}

	if exist, ip, e := service.Vps.NameExist(f.Name); e != nil {
		controllers.CallBack.Error(c, 3)
		return
	} else if exist && ip != c.ClientIP() {
		controllers.CallBack.Error(c, 4)
		return
	} else if e = service.Vps.SetName(c.ClientIP(), f.Name); e != nil {
		controllers.CallBack.Error(c, 3)
		return
	}

	modules.Online.Up(c.ClientIP())

	controllers.CallBack.Default(c)
}
