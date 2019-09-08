package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/app/v1/service_project"
	"github.com/mangmang/app/v1/service_role"
	"github.com/mangmang/app/v1/service_user"
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
		apiV1.POST("/user/changePassWord", service_user.ChangePassWord)
		apiV1.GET("/user/searchUser", service_user.SearchUser)
		apiV1.GET("/user/userInfo", service_user.GetUserInfo)

		apiV1.GET("/user/businessCard", service_user.GetBusinessCard)
		apiV1.GET("/user/businessCard/:key", service_user.GetBusinessCard)
		apiV1.POST("/user/businessCard", service_user.CreateBusinessCard)
		apiV1.PUT("/user/businessCard/:key", service_user.UpdateBusinessCard)
		apiV1.DELETE("/user/businessCard/:key", service_user.DeleteBusinessCard)

		apiV1.GET("/role", service_role.GetRole)
		apiV1.GET("/role/:key", service_role.GetRole)
		apiV1.POST("/role", service_role.CreateRole)
		apiV1.PUT("/role/:key", service_role.UpadteRole)

		apiV1.GET("/project", service_project.GetProject)
		apiV1.GET("/project/:key", service_project.GetProject)
		apiV1.POST("/project", service_project.CreateProject)
		apiV1.PUT("/project/:key", service_project.UpdateProject)

		apiV1.GET("/task", service_project.GetTasks)
		apiV1.GET("/task/:key", service_project.GetTasks)
		apiV1.POST("/task", service_project.CreateTask)
		apiV1.DELETE("/task/:key", service_project.DeleteTask)

		apiV1.GET("/fatherTask", service_project.GetFatherTask)

		apiV1.GET("/projectUser", service_project.GetProjectUser)
		apiV1.POST("/projectUser", service_project.AddProjectUser)
		apiV1.PUT("/projectUser", service_project.ChangeProjectUserRole)
		apiV1.DELETE("/projectUser", service_project.RemoveProjectUser)

	}

	return r
}
