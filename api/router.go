package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/api/v1/service_user"
	"github.com/mangmang/api/web"
	"github.com/mangmang/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.LoadHTMLGlob("templates/index.html")

	r.GET("/", web.Home)
	r.GET("/home", web.Home)

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/test", service_user.Test)
		apiV1.GET("/user/getVerificationCode", service_user.GetVerificationCode)
		apiV1.POST("/user/phoneRegister", service_user.PhoneRegister)
		apiV1.POST("/user/loginAPW", service_user.UserLoginAPW)
		apiV1.POST("/user/loginAC", service_user.UserLoginAC)
		apiV1.PUT("/user/changeUserInfo", service_user.ChangeUserInfo)
		apiV1.POST("/user/uploadAvatar", service_user.UploadAvatar)
		apiV1.POST("/user/changePassWord", service_user.ChangePassWord)
		apiV1.GET("/user/searchUser", service_user.SearchUser)

		apiV1.GET("/user/businessCard", service_user.GetBusinessCard)
		apiV1.GET("/user/businessCard/:key", service_user.GetBusinessCard)
		apiV1.POST("/user/businessCard", service_user.CreateBusinessCard)
		apiV1.PUT("/user/businessCard", service_user.UpdateBusinessCard)
		apiV1.DELETE("/user/businessCard/:key", service_user.DeleteBusinessCard)
	}

	return r
}
