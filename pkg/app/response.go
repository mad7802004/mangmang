package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/e"
)

type Gin struct {
	C *gin.Context
	M *e.Msg
}

func New(c *gin.Context) *Gin {
	newMsg := Gin{c, e.NewMsg(c.GetHeader("lang"))}
	return &newMsg
}

func (g *Gin) AddField(key string, value interface{}) {
	g.M.AddField(key, value)
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.M.Update(errCode, data)
	msg := g.M.M
	delete(msg, "lang")
	g.C.JSON(httpCode, msg)
	return
}
