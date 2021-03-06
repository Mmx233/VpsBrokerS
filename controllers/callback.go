package controllers

import (
	"github.com/gin-gonic/gin"
)

type call struct{}

var CallBack call

type CallBackMessage struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
}

func (*call) e(c *gin.Context, code uint, httpCode int) {
	c.AsciiJSON(httpCode, CallBackMessage{
		code,
		[]string{},
	})
	c.Abort()
}
func (a *call) Error(c *gin.Context, code uint) {
	a.e(c, code, 400)
}

func (*call) Success(c *gin.Context, data interface{}) {
	c.AsciiJSON(200, CallBackMessage{
		0,
		data,
	})
}

func (a *call) Default(c *gin.Context) {
	a.Success(c, []string{})
}
