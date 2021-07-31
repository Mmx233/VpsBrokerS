package middlewares

import (
	"github.com/Mmx233/VpsBrokerS/global"
	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != global.Config.Settings.AccessKey {
			c.Abort()
		}
	}
}
