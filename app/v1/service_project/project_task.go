package service_project

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 获取项目任务
func GetTasks(c *gin.Context) {
	appG := app.New(c)
	key := c.Param("key")
	projectId := c.Query("project_id")

	// 项目未填
	if projectId == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 判断是否查询一个任务信息
	if key != "" {
		query, err := models.FindTask(key)
		if err != nil {
			appG.Response(http.StatusOK, e.NoResourcesFound, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, query)
		return
	}

	// 获取分页
	page, size := app.GetPageSize(c)
	tasks, total, err := models.FindProjectIdTasks(projectId, page, size)
	// 未找到数据
	if err != nil || len(tasks) == 0 {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}
	data := map[string]interface{}{
		"tasks": tasks,
		"total": total,
		"size":  size,
		"page":  page,
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}

// 创建项目任务
func CreateTask(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 更新项目任务
func UpdateTask(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 删除项目任务
func DeleteTask(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}


// 项目任务指派
func DistributionTask(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
