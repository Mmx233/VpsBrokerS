package main

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerS/global"
	"github.com/Mmx233/VpsBrokerS/router"
	"log"
)

func main() {
	if e := router.G.Run(":" + fmt.Sprint(global.Config.Settings.Port)); e != nil {
		log.Fatalln(e)
	}
}
