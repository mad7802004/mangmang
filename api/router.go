package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/setting"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "成功"})
		})

	}

	return r
}
