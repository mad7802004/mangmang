package models

import (
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
)

type Project struct {
	ProjectId      string         `json:"project_id"gorm:"primary_key"` // 项目ID
	ProjectName    string         `json:"project_name"`                 // 项目名称
	ProjectContent string         `json:"project_content"`              // 项目描述内容
	CreateTime     utils.JSONTime `json:"create_time"`                  // 创建日期
	UpdateTime     utils.JSONTime `json:"update_time"`                  // 更新日期
	DataStatus     int8           `json:"data_status"gorm:"default"`    // 状态
}

type UserProjectMapping struct {
	Id         string         `json:"id"gorm:"primary_key"`      // 项目管理ID
	UserId     string         `json:"user_id"`                   // 用户ID
	ProjectId  string         `json:"project_id"`                // 项目ID
	RoleId     string         `json:"role_id"`                   // 角色ID
	CreateTime utils.JSONTime `json:"create_time"`               // 创建时间
	UpdateTime utils.JSONTime `json:"update_time"`               // 更新时间
	DataStatus int8           `json:"data_status"gorm:"default"` // 状态
}

// 根据项目ID查询项目信息
func FindProject(projectId string) (*Project, error) {
	var info Project

	err := Orm.Model(&Project{}).Where("project_id = ?", projectId).Find(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil

}

// 根据用户ID查询用户所以项目
func FindUserProject(userId string, page, size int) ([]*Project, int, error) {
	var project []*Project
	var total int
	query := Orm.Model(&Project{}).
		Joins("inner join user_project_mapping on user_project_mapping.project_id=project.project_id ").
		Where("user_project_mapping.user_id = ? and data_status=?", userId, e.Enable)

	err := query.Offset((page - 1) * size).Limit(size).
		Find(&project).Error
	if err != nil || len(project) == 0 {
		return nil, 0, err
	}
	query.Count(&total)
	return project, total, nil
}

//