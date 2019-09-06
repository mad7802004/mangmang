package service_project

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
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
	var obj struct {
		FatherTaskId string          `json:"father_task_id"binding:"omitempty,uuid4"`
		ProjectId    string          `json:"project_id"binding:"required,uuid4"`
		UserId       string          `json:"user_id"binding:"omitempty,uuid4"`
		TaskName     string          `json:"task_name"binding:"required,max=50"`
		TaskPriority int             `json:"task_priority"binding:"gte=0,lte=6"`
		TaskType     int             `json:"task_type"binding:"gte=0,lte=3"`
		TaskStatus   int             `json:"task_status"binding:"gte=0,lte=3"`
		TaskContent  string          `json:"task_content"binding:"required"`
		StartTime    *utils.JSONTime `json:"start_time"binding:"omitempty"`
		EndTime      *utils.JSONTime `json:"end_time"binding:"omitempty"`
	}
	appG := app.New(c)
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	if !models.IsExistProject(obj.ProjectId) {
		appG.Response(http.StatusOK, e.ProjectDoesNotExist, nil)
		return
	}

	if obj.FatherTaskId != "" && !models.IsExistTask(obj.FatherTaskId) {
		appG.Response(http.StatusOK, e.FatherTaskDoesNotExist, nil)
		return
	}

	if obj.UserId != "" && !models.IsExistProjectUser(obj.ProjectId, obj.UserId) {
		appG.Response(http.StatusOK, e.ProjectUserDoesNotExist, nil)
		return
	}

	taskNumber := models.FindMaxTaskNumber(obj.ProjectId)

	newTask := &models.Task{
		TaskId:       utils.GetUUID(),
		FatherTaskId: obj.FatherTaskId,
		ProjectId:    obj.ProjectId,
		UserId:       obj.UserId,
		TaskNumber:   taskNumber + 1,
		TaskName:     obj.TaskName,
		TaskPriority: obj.TaskPriority,
		TaskType:     obj.TaskType,
		TaskContent:  obj.TaskContent,
		TaskSchedule: 0,
		TaskStatus:   obj.TaskStatus,
		StartTime:    obj.StartTime,
		EndTime:      obj.EndTime,
	}

	// 创建任务
	if !models.Create(newTask) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

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

// 获取任务列表，用于指定父级任务
func GetFatherTask(c *gin.Context) {
	appG := app.New(c)
	projectId := c.Query("project_id")
	// 项目未填
	if projectId == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	fatherTasks, err := models.FindFatherTasks(projectId)
	if err != nil {
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, fatherTasks)
	return

}

// 项目任务指派
func DistributionTask(c *gin.Context) {
	appG := app.New(c)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
