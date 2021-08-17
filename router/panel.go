package router

import (
	controllers "github.com/Mmx233/VpsBrokerS/controllers/panel"
	"github.com/gin-gonic/gin"
)

func routerPanel(G *gin.RouterGroup) {
	G.GET("/vps", controllers.Client.List)
}
