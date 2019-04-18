package models

import "github.com/mangmang/pkg/utils"

type UserProjectMapping struct {
	Id         string         `json:"id"gorm:"primary_key"`      // 项目管理ID
	UserId     string         `json:"user_id"`                   // 用户ID
	ProjectId  string         `json:"project_id"`                // 项目ID
	RoleId     string         `json:"role_id"`                   // 角色ID
	CreateTime utils.JSONTime `json:"create_time"`               // 创建时间
	UpdateTime utils.JSONTime `json:"update_time"`               // 更新时间
	DataStatus int8           `json:"data_status"gorm:"default"` // 状态
}
