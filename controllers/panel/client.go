package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/VpsBrokerS/models/form"
	"github.com/Mmx233/VpsBrokerS/service"
	"github.com/gin-gonic/gin"
)

type client struct{}

var Client client

func (*client) List(c *gin.Context) {
	d, e := service.Vps.List()
	if e != nil {
		controllers.CallBack.Error(c, 3)
		return
	}

	controllers.CallBack.Success(c, d)
}

func (*client) Del(c *gin.Context) {
	var f form.DelVps
	if e := c.ShouldBind(&f); e != nil {
		controllers.CallBack.Error(c, 2)
		return
	}

	if !service.Vps.IdExist(f.ID) {
		controllers.CallBack.Error(c, 5)
		return
	}

	if service.Vps.Del(f.ID) != nil {
		controllers.CallBack.Error(c, 3)
		return
	}

	controllers.CallBack.Default(c)
}
