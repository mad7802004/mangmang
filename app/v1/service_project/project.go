package service_project

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
	"net/http"
)

// 获取项目列表或项目
func GetProject(c *gin.Context) {
	appG := app.New(c)
	key := c.Param("key")
	userId := c.Query("user_id")

	// 用户未填
	if userId == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 判断是否查询一个项目信息
	if key != "" {
		query, err := models.FindProject(key)
		if err != nil {
			appG.Response(http.StatusOK, e.NoResourcesFound, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, query)
		return
	}

	// 获取分页
	page, size := app.GetPageSize(c)
	data, total, err := models.FindUserProject(userId, page, size)
	// 未找到数据
	if err != nil || len(data) == 0 {
		appG.AddField("total", total)
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	appG.AddField("total", total)
	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}

// 新建项目
func CreateProject(c *gin.Context) {

	var obj struct {
		UserId         string `json:"user_id"binding:"uuid4"`
		ProjectName    string `json:"project_name"binding:"max=20"`
		ProjectContent string `json:"project_content"`
	}

	appG := app.New(c)
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}
	// 判断用户是否存在
	if !models.IsExistUser(obj.UserId) {
		appG.Response(http.StatusOK, e.AccountDoesNotExist, nil)
		return
	}
	projectId := utils.GetUUID()
	newProject := &models.Project{
		ProjectId:      projectId,
		ProjectName:    obj.ProjectName,
		ProjectContent: obj.ProjectContent,
	}
	newUserProjectMapping := &models.UserProjectMapping{
		UserId:    obj.UserId,
		ProjectId: projectId,
		RoleId:    "1", //TODO:未完成角色绑定
	}
	if !models.Create(newProject, newUserProjectMapping) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

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
