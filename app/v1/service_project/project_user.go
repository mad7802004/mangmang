package service_project

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 获取项目成员列表
func GetProjectUser(c *gin.Context) {

	appG := app.New(c)
	projectId := c.Query("project_id")

	userList, err := models.FindProjectIdUser(projectId)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, userList)
	return
}

// 添加项目成员
func AddProjectUser(c *gin.Context) {
	var obj struct {
		UserId    string `json:"user_id"binding:"uuid4"`
		RoleId    string `json:"role_id"binding:"uuid4"`
		ProjectId string `json:"project_id"binding:"uuid4"`
	}

	appG := app.New(c)

	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 判断项目是否存在
	if !models.IsExistProject(obj.ProjectId) {
		appG.Response(http.StatusOK, e.ProjectDoesNotExist, nil)
		return
	}

	// 判断用户是否存在
	if !models.IsExistUser(obj.UserId) {
		appG.Response(http.StatusOK, e.AccountDoesNotExist, nil)
		return
	}

	// 判断项目下用户是否存在
	if !models.IsExistProjectUser(obj.ProjectId, obj.UserId) {
		appG.Response(http.StatusOK, e.ProjectUserExist, nil)
		return
	}

	// 判断角色是否存在
	if !models.IsExistRole(obj.RoleId) {
		appG.Response(http.StatusOK, e.RoleDoesNotExist, nil)
		return
	}

	newUserProjectMapping := &models.UserProjectMapping{
		UserId:    obj.UserId,
		ProjectId: obj.ProjectId,
		RoleId:    obj.RoleId,
	}

	// 创建项目成员
	if !models.Create(newUserProjectMapping) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 修改项目成员权限
func ChangeProjectUserRole(c *gin.Context) {
	var obj struct {
		MappingId string `json:"mapping_id"binding:"uuid4"`
		UserId    string `json:"user_id"binding:"uuid4"`
		ProjectId string `json:"project_id"binding:"uuid4"`
		RoleId    string `json:"role_id"binding:"uuid4"`
	}

	appG := app.New(c)

	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 判断角色是否存在
	if !models.IsExistRole(obj.RoleId) {
		appG.Response(http.StatusOK, e.RoleDoesNotExist, nil)
		return
	}
	// 查询查询成员权限关联
	mapping, err := models.FindProjectUserMapping(obj.MappingId, obj.UserId, obj.ProjectId)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	// 更新用户权限
	if !models.UpdateProjectUserMapping(mapping, obj) {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 移除项目成员 TODO:权限未验证且项目成员能删除自己、项目中一个成员都不剩
func RemoveProjectUser(c *gin.Context) {
	var obj struct {
		MappingId string `json:"mapping_id"binding:"uuid4"`
		UserId    string `json:"user_id"binding:"uuid4"`
		ProjectId string `json:"project_id"binding:"uuid4"`
	}

	appG := app.New(c)

	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 查询查询成员权限关联
	mapping, err := models.FindProjectUserMapping(obj.MappingId, obj.UserId, obj.ProjectId)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	if !models.DeleteProjectUserMapping(mapping) {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}
