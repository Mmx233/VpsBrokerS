package controllers

import (
	"github.com/Mmx233/VpsBrokerS/controllers"
	"github.com/Mmx233/VpsBrokerS/models/form"
	"github.com/Mmx233/VpsBrokerS/service"
	"github.com/Mmx233/VpsBrokerS/service/modules"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upper = websocket.Upgrader{
	HandshakeTimeout: time.Minute,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Init(c *gin.Context) {
	var f form.VpsInit
	if e := c.ShouldBind(&f); e != nil {
		controllers.CallBack.Error(c, 2)
		return
	}

	if exist, ip, e := service.Vps.NameExist(f.Name); e != nil {
		controllers.CallBack.Error(c, 3)
		return
	} else if exist && ip != c.ClientIP() {
		controllers.CallBack.Error(c, 4)
		return
	} else if e = service.Vps.SetName(c.ClientIP(), f.Name); e != nil {
		controllers.CallBack.Error(c, 3)
		return
	}

	conn, err := upper.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// bad request
		return
	}

	ip := c.ClientIP()

	modules.Pool.Add(ip, conn, f.Port)

	go func() {
		defer func() {
			recover()
		}()

		for {
			var data form.HeartBeat

			err := conn.ReadJSON(&data)
			if err != nil {
				// disconnected or timeout
				_ = conn.Close()
				modules.Pool.Lose(ip)
				return
			}

			switch data.Type {
			case "up":
				_ = service.Event.Up(data.TargetIp, data.Time)
			case "down":
				_ = service.Event.Down(data.TargetIp, data.Time)
			default:
				continue
			}
		}
	}()
}
