package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/VpsBrokerS/service/modules"
	"github.com/gin-gonic/gin"
)

// GetSelf 客户端获取自身ip接口
func GetSelf(c *gin.Context) {
	controllers.CallBack.Success(c, map[string]interface{}{
		"ip": c.ClientIP(),
	})
}

// GetListInfo 客户端主动获取客户端列表接口
func GetListInfo(c *gin.Context) {
	controllers.CallBack.Success(c, modules.Pool.GetListInfo())
}
