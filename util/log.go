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

func (a *event) Up(name string) {
	a.e("up", name+" 在线率上升")
}

func (a *event) Down(name string) {
	a.e("down", name+" 在线率下降")
}
