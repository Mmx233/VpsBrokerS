package router

import (
	"github.com/Mmx233/VpsBrokerS/middlewares"
	"github.com/Mmx233/secure"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	G = gin.Default()

	G.Use(secure.Main(), middlewares.Auth())
}
