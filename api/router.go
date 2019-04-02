package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/api/v1/service_user"
	"github.com/mangmang/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/test", service_user.Test)
	}

	return r
}
