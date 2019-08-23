package models

import "github.com/mangmang/pkg/utils"

type Task struct {
	TaskId       string         `json:"task_id"`                     // 任务ID
	FatherTaskId string         `json:"father_task_id"`              // 父级任务
	ProjectId    string         `json:"project_id"`                  // 项目ID
	UserId       string         `json:"user_id"`                     // 任务人员
	TaskName     string         `json:"task_name"`                   // 任务名称
	TaskPriority int            `json:"task_priority"gorm:"default"` // 任务优先级
	TaskType     string         `json:"task_type"`                   // 任务类型
	TaskContent  string         `json:"task_content"`                // 任务内容
	TaskSchedule int            `json:"task_schedule"gorm:"default"` // 任务进度
	TaskStatus   string         `json:"task_status"gorm:"default"`   // 任务状态
	StartTime    utils.JSONTime `json:"start_time"`                  // 任务开始时间
	EndTime      utils.JSONTime `json:"end_time"`                    // 任务预计完成时间
	CreateTime   utils.JSONTime `json:"create_time"`                 // 创建时间
	UpdateTime   utils.JSONTime `json:"update_time"`                 // 更新时间
	DataStatus   int8           `json:"data_status"gorm:"default"`   // 状态
}

// 根据任务ID查询任务信息
func FindTask(taskId string) (*Task, error) {
	var info Task

	err := Orm.Model(&Task{}).Where("task_id = ?", taskId).Find(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil

}

// 根据项目ID查询所有任务
func FindProjectIdTasks(projectId string, page, size int) ([]*Task, int, error) {
	var tasks []*Task
	var total int
	query := Orm.Model(&Task{}).Where("project_id = ?", projectId).Order("create_time", false)

	err := query.Offset((page - 1) * size).Limit(size).
		Find(&tasks).Error
	if err != nil || len(tasks) == 0 {
		return nil, 0, err
	}
	query.Count(&total)
	return tasks, total, nil
}
