package util

import (
	"fmt"
	"log"
)

type event struct{}

var Event event

func (*event) e(name string, content string) {
	log.Println(fmt.Sprintf("[%s] %s", name, content))
}

func (a *event) NewClientConn(name string) {
	a.e("client", name+" 已连接")
}

func (a *event) LostClient(name string) {
	a.e("client", name+" 连接断开")
}
