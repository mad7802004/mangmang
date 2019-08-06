package service_role

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 获取角色列表或角色
func GetRole(c *gin.Context) {
	appG := app.New(c)
	key := c.Param("key")
	// 判断是否查询一个角色
	if key != "" {
		query, err := models.FindRole(key)
		if err != nil {
			appG.Response(http.StatusOK, e.NoResourcesFound, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, query)
		return
	}
	// 获取分页
	page, size := app.GetPageSize(c)
	roles, total, err := models.FindRoleList(page, size)
	// 未找到数据
	if err != nil || len(roles) == 0 {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}
	data := map[string]interface{}{
		"roles": roles,
		"total": total,
		"size":  size,
		"page":  page,
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
	return

}

// 创建角色并指定权限
func CreateRole(c *gin.Context) {
	var obj struct {
		RoleLevel int    `json:"role_level"binding:"required"`
		RoleName  string `json:"role_name"binding:"required,max=50"`
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	newRole := &models.Role{
		RoleLevel: obj.RoleLevel,
		RoleName:  obj.RoleName,
	}
	// 创建新权限
	if !models.Create(newRole) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 修改角色、权限
func UpadteRole(c *gin.Context) {
	var obj struct {
		RoleLevel int    `json:"role_level"binding:"required"`
		RoleName  string `json:"role_name"binding:"required,max=50"`
	}

	appG := app.New(c)
	key := c.Param("key")

	//参数解析失败或者键值为空
	if c.ShouldBindJSON(&obj) != nil || key == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 查询角色
	role, err := models.FindRole(key)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	// 更新信息
	if !models.UpdateRole(role, obj) {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 删除角色
func DeleteRole(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
