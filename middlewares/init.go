package middlewares

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/secure"
	"github.com/gin-gonic/gin"
)

func init() {
	secure.Init(func(c *gin.Context) {
		controllers.CallBack.Error(c, 1)
	})
}
