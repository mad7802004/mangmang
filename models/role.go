package models

import "github.com/mangmang/pkg/utils"

type Role struct {
	RoleId     string         `json:"role_id"gorm:"primary_key"` // 角色ID
	RoleLevel  int            `json:"role_level"`                // 角色权限
	RoleName   string         `json:"role_name"`                 // 角色名称
	CreateTime utils.JSONTime `json:"-"`                         // 创建时间
	UpdateTime utils.JSONTime `json:"-"`                         // 更新时间
	DataStatus int8           `json:"-"gorm:"default"`           // 状态
}
