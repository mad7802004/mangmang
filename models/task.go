package models

import "github.com/mangmang/pkg/utils"

type Task struct {
	TaskId                string         `json:"task_id"`                     // 任务ID
	FatherTaskId          string         `json:"father_task_id"`              // 父级任务
	ProjectId             string         `json:"project_id"`                  // 项目ID
	TaskName              string         `json:"task_name"`                   // 任务名称
	TaskType              string         `json:"task_type"`                   // 任务类型
	TaskContent           string         `json:"task_content"`                // 任务内容
	TaskSchedule          int            `json:"task_schedule"gorm:"default"` // 任务进度
	TaskStatus            string         `json:"task_status"gorm:"default"`   // 任务状态
	StartingTime          utils.JSONTime `json:"starting_time"`               // 任务开始时间
	PlannedCompletionTime utils.JSONTime `json:"planned_completion_time"`     // 任务预计完成时间
	CreateTime            utils.JSONTime `json:"create_time"`                 // 创建时间
	UpdateTime            utils.JSONTime `json:"update_time"`                 // 更新时间
	DataStatus            int8           `json:"data_status"gorm:"default"`   // 状态
}
