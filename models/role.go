package models

import (
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
)

type Role struct {
	RoleId     string         `json:"role_id"gorm:"primary_key"` // 角色ID
	RoleLevel  int            `json:"role_level"`                // 角色权限
	RoleName   string         `json:"role_name"`                 // 角色名称
	CreateTime utils.JSONTime `json:"-"`                         // 创建时间
	UpdateTime utils.JSONTime `json:"-"`                         // 更新时间
	DataStatus int8           `json:"-"gorm:"default"`           // 状态
}

// 根据角色ID查询角色
func FindRole(roleId string) (*Role, error) {
	var info Role

	err := Orm.Model(&Role{}).Where("role_id = ? ", roleId).Find(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// 判断角色是否存在
func IsExistRole(roleId string) bool {

	_, err := FindRole(roleId)
	if err != nil {
		return false
	}
	return true
}

// 查询角色列表
func FindRoleList(page, size int) ([]*Role, int, error) {
	var role []*Role
	var total int

	query := Orm.Model(&Role{}).Where(" data_status=?", e.Enable)
	err := query.Offset((page - 1) * size).Limit(size).Find(&role).Error
	if err != nil || len(role) == 0 {
		return nil, 0, err
	}
	query.Count(&total)

	return role, total, nil
}

// 更新角色信息
func UpdateRole(role *Role, data interface{}) bool {
	err := Orm.Model(&role).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}
