package router

import (
	controllers "github.com/Mmx233/VpsBrokerS/controllers/client"
	"github.com/gin-gonic/gin"
)

func routerClient(G *gin.RouterGroup) {
	G.GET("/", controllers.Init)
	G.GET("/self", controllers.GetSelf)
	G.GET("/list", controllers.GetListInfo)
}
