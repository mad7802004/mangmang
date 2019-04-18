package service_project

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 获取项目列表
func GetProjects(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 新建项目
func CreateProject(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 更新项目
func UpdateProject(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 删除项目
func DeleteProject(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 添加项目成员
func AddProjectUser(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 修改项目成员权限
func ChangeProjectUserRole(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 移除项目成员
func RemoveProjectUser(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 项目进度
func ProjectProgress(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}
