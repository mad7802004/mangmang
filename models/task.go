package models

import (
	"github.com/mangmang/pkg/utils"
)

type Task struct {
	TaskId         string          `json:"task_id"gorm:"primary_key"`   // 任务ID
	FatherTaskId   string          `json:"father_task_id"`              // 父级任务
	ProjectId      string          `json:"project_id"`                  // 项目ID
	TaskFinisherId string          `json:"task_finisher_id"`            // 任务人员
	TaskCreatorId  string          `json:"task_creator_id"`             // 任务创建人
	TaskNumber     int             `json:"task_number"`                 // 任务编码
	TaskName       string          `json:"task_name"`                   // 任务名称
	TaskPriority   *int            `json:"task_priority"gorm:"default"` // 任务优先级
	TaskType       *int            `json:"task_type"`                   // 任务类型
	TaskContent    string          `json:"task_content"`                // 任务内容
	TaskSchedule   *int            `json:"task_schedule"gorm:"default"` // 任务进度
	TaskStatus     *int            `json:"task_status"gorm:"default"`   // 任务状态
	StartTime      *utils.JSONDate `json:"start_time"gorm:"default"`    // 任务开始时间
	EndTime        *utils.JSONDate `json:"end_time"`                    // 任务预计完成时间
	CreateTime     utils.JSONTime  `json:"create_time"`                 // 创建时间
	UpdateTime     utils.JSONTime  `json:"update_time"`                 // 更新时间
	DataStatus     int8            `json:"data_status"gorm:"default"`   // 状态
}

type ChildTask struct {
	Task
	TaskFinisher string `json:"task_finisher"`
}

// 根据任务ID查询子任务列表
func FindChildTask(fatherTaskId string) ([]*ChildTask, error) {
	var childTaskList []*ChildTask
	err := Orm.Model(&Task{}).
		Select("task.*,user.name as task_finisher").
		Joins("left join user on user.user_id = task.task_finisher_id").
		Where("father_task_id = ?", fatherTaskId).
		Scan(&childTaskList).Order("task_number").Error
	if err != nil {
		return nil, err
	}
	return childTaskList, nil
}

type QueryTaskInfo struct {
	Task
	TaskFinisher     string `json:"task_finisher"`
	TaskCreator      string `json:"task_creator"`
	CreatorAvatar    string `json:"creator_avatar"`
	FatherTaskId     string `json:"father_task_id"`
	FatherTaskNumber int    `json:"father_task_number"`
	FatherTaskName   string `json:"father_task_name"`
}

// 根据任务ID查询任务信息，并查询父级任务和创建人和被指派人信息
func FindTaskInfo(taskId string) (*QueryTaskInfo, error) {
	var info QueryTaskInfo

	err := Orm.Model(&Task{}).Select("task.*,"+
		"creator.name as task_creator, creator.avatar_url as creator_avatar,"+
		"finisher.name as task_finisher, father.task_id as  father_task_id,"+
		"father.task_number as father_task_number,father.task_name as father_task_name").
		Joins("left join task as father on father.task_id = task.task_finisher_id").
		Joins("inner join user as creator on creator.user_id=task.task_creator_id").
		Joins("left join user as finisher on finisher.user_id = task.task_finisher_id ").
		Where("task.task_id = ?", taskId).
		Scan(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil

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

// 判断任务是否存在
func IsExistTask(taskId string) bool {
	_, err := FindTask(taskId)
	if err != nil {
		return false
	}
	return true

}

type QueryProjectTaskList struct {
	Task
	UserName string `json:"user_name"`
}

// 根据项目ID查询所有任务
func FindProjectIdTasks(projectId string, page, size int) ([]*QueryProjectTaskList, int, error) {

	var tasks []*QueryProjectTaskList
	var total int
	query := Orm.Model(&Task{}).Select("task.*,user.name as user_name").
		Joins("left join user on user.user_id = task.task_finisher_id").
		Where("project_id = ?", projectId).Order("task_number DESC", false)

	err := query.Offset((page - 1) * size).Limit(size).
		Scan(&tasks).Error
	if err != nil || len(tasks) == 0 {
		return nil, 0, err
	}
	query.Count(&total)
	return tasks, total, nil
}

// 根据项目ID查询当前项目任务最大编号
func FindMaxTaskNumber(projectId string, ) int {
	var result struct {
		MaxNumber int `json:"max_number"`
	}

	query := Orm.Model(&Task{}).Select("max(task_number) as max_number").Where("project_id = ?", projectId)
	err := query.Scan(&result).Error
	if err != nil {
		return 0
	}
	return result.MaxNumber

}

type FatherTasks struct {
	TaskId     string `json:"task_id"`
	TaskNumber int    `json:"task_number"`
	TaskName   string `json:"task_name"`
}

// 根据项目ID获取任务列表，用于关联父任务
func FindFatherTasks(projectId string) ([]*FatherTasks, error) {
	var fatherTasks []*FatherTasks

	query := Orm.Model(&Task{}).Select("task_id,task_number,task_name").
		Where("project_id = ?", projectId).Order("task_number", false)

	err := query.Scan(&fatherTasks).Error
	if err != nil {
		return nil, err
	}
	return fatherTasks, nil

}

// 更新用户信息
func UpdateTask(task *Task, data interface{}) bool {
	err := Orm.Model(&task).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 删除任务
func DeleteTask(task *Task) bool {
	err := Orm.Delete(&task).Error
	if err != nil {
		return false
	}
	return true
}
