package models

import "github.com/mangmang/pkg/utils"

type Project struct {
	ProjectId      string         `json:"project_id"gorm:"primary_key"` // 项目ID
	ProjectName    string         `json:"project_name"`                 // 项目名称
	ProjectContent string         `json:"project_content"`              // 项目描述内容
	CreateTime     utils.JSONTime `json:"create_time"`                  // 创建日期
	UpdateTime     utils.JSONTime `json:"update_time"`                  // 更新日期
	DataStatus     int8           `json:"data_status"gorm:"default"`    // 状态
}
