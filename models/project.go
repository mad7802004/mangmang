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

// 判断项目是否存在
func IsExistProject(projectId string) bool {

	_, err := FindProject(projectId)
	if err != nil {
		return false
	}
	return true

}

// 根据用户ID查询用户所有项目
func FindUserProject(userId string, page, size int) ([]*Project, int, error) {
	var project []*Project
	var total int
	query := Orm.Model(&Project{}).
		Joins("inner join user_project_mapping on user_project_mapping.project_id=project.project_id ").
		Where("user_project_mapping.user_id = ?", userId)

	err := query.Offset((page - 1) * size).Limit(size).
		Find(&project).Error
	if err != nil || len(project) == 0 {
		return nil, 0, err
	}
	query.Count(&total)
	return project, total, nil
}

// 更新项目信息
func UpdateProject(project *Project, data interface{}) bool {
	err := Orm.Model(&project).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 根据项目ID查询所有的项目成员
func FindProjectIdUser(projectId string) (interface{}, error) {
	var userList []struct {
		UserName  string `json:"user_name"`
		UserId    string `json:"user_id"`
		RoleId    string `json:"role_id"`
		RoleName  string `json:"role_name"`
		ProjectId string `json:"project_id"`
	}
	err := Orm.Model(&UserProjectMapping{}).
		Select("user.name as user_name,user.user_id as user_id," +
			"role.role_id as role_id,role.role_name as role_name,user_project_mapping.project_id as project_id ").
		Joins("inner join user on user.user_id = user_project_mapping.user_id ").
		Joins("inner join role on role.role_id = user_project_mapping.role_id ").
		Where("user_project_mapping.project_id = ? and user_project_mapping.data_status=?", projectId, e.Enable).
		Scan(&userList).Error
	if err != nil || len(userList) == 0 {
		return nil, err
	}
	return userList, nil

}

// 判断项目用户是否存在
func IsExistProjectUser(projectId, userId string) bool {
	var total int
	err := Orm.Model(&UserProjectMapping{}).
		Where("project_id =? and user_id=?", projectId, userId).Count(&total).Error
	if err != nil || total > 0 {
		return false
	}

	return true
}

// 查询查询成员权限关联
func FindProjectUserMapping(mappingId, userId, projectId string) (*UserProjectMapping, error) {
	var mapping UserProjectMapping

	err := Orm.Model(&UserProjectMapping{}).
		Where("id = ? and user_id =? and project_id=? ", mappingId, userId, projectId).Find(&mapping).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil

}

// 更新用户关联信息
func UpdateProjectUserMapping(userProjectMapping *UserProjectMapping, data interface{}) bool {
	err := Orm.Model(&userProjectMapping).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}


// 移除项目用户
func DeleteProjectUserMapping(userProjectMapping *UserProjectMapping) bool {
	err := Orm.Delete(&userProjectMapping).Error
	if err != nil {
		return false
	}
	return true
}
