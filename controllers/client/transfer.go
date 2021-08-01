package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/VpsBrokerS/service/modules"
	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	controllers.CallBack.Success(c, map[string]interface{}{
		"ip": c.ClientIP(),
	})
}

func GetListInfo(c *gin.Context) {
	controllers.CallBack.Success(c, modules.Pool.GetListInfo())
}
