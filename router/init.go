package router

import (
	"github.com/Mmx233/VpsBrokerS/middlewares"
	"github.com/Mmx233/secure"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	G = gin.New()
	G.Use(gin.Recovery(), secure.Main(), middlewares.Auth())

	routerClient(G.Group("/c"))
	routerPanel(G.Group("/p"))
}
