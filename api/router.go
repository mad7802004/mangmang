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
		apiV1.GET("/user/getVerificationCode", service_user.GetVerificationCode)
		apiV1.POST("/user/phoneRegister", service_user.PhoneRegister)
		apiV1.POST("/user/loginAPW", service_user.UserLoginAPW)
		apiV1.POST("/user/loginAC", service_user.UserLoginAC)
		apiV1.PUT("/user/changeUserInfo", service_user.ChangeUserInfo)
		apiV1.POST("/user/uploadAvatar", service_user.UploadAvatar)
	}

	return r
}
