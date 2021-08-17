package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
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

}
