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
		UserId         string `json:"user_id"binding:"required,uuid4"`
		ProjectName    string `json:"project_name"binding:"required,max=20"`
		ProjectContent string `json:"project_content"`
		RoleId         string `json:"role_id"binding:"required,uuid4"`
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

	// 判断角色是否存在
	if !models.IsExistRole(obj.RoleId) {
		appG.Response(http.StatusOK, e.RoleDoesNotExist, nil)
		return
	}

	// 新建项目
	projectId := utils.GetUUID()
	newProject := &models.Project{
		ProjectId:      projectId,
		ProjectName:    obj.ProjectName,
		ProjectContent: obj.ProjectContent,
	}

	newUserProjectMapping := &models.UserProjectMapping{
		UserId:    obj.UserId,
		ProjectId: projectId,
		RoleId:    obj.RoleId,
	}

	// 创建项目
	if !models.Create(newProject, newUserProjectMapping) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 更新项目
func UpdateProject(c *gin.Context) {
	var obj struct {
		ProjectName    string `json:"project_name"binding:"required,max=20"`
		ProjectContent string `json:"project_content"`
	}
	appG := app.New(c)
	key := c.Param("key")

	//参数解析失败或者键值为空
	if c.ShouldBindJSON(&obj) != nil || key == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}


	// 查询项目是否存在
	project, err := models.FindProject(key)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	// 更新项目信息
	if !models.UpdateProject(project, obj) {
		appG.Response(http.StatusOK, e.UpdateFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}

// 删除项目 TODO:权限未验证，只有admin才能删除项目且移除所有项目成员
func DeleteProject(c *gin.Context) {
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
