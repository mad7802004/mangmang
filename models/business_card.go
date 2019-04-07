package models

import (
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
)

// 名片表
type BusinessCard struct {
	Id         string         `json:"id"gorm:"primary_key"` // id
	UserId     string         `json:"user_id"`              // 用户ID
	Name       string         `json:"name"`                 // 姓名
	Company    string         `json:"company"`              // 公司
	Position   string         `json:"position"`             // 职位
	Phone      string         `json:"phone"`                // 电话
	Qq         string         `json:"qq"`                   // QQ号
	Wx         string         `json:"wx"`                   // 微信号
	CreateTime utils.JSONTime `json:"create_time"`          // 创建时间
	UpdateTime utils.JSONTime `json:"update_time"`          // 更新时间
	DataStatus int8           `json:"-"gorm:"default"`      // 0删除，1有效
}

// 根据用户ID查询用户名片
func FindUserBusinessCard(userId string, page, size int) ([]*BusinessCard, int, error) {
	var businessCard []*BusinessCard
	var total int
	query := Orm.Model(&BusinessCard{}).
		Where("user_id = ? and data_status=?", userId, e.Enable)

	err := query.Offset((page - 1) * size).Limit(size).
		Find(&businessCard).Error
	if err != nil || len(businessCard) == 0 {
		return nil, 0, err
	}
	query.Count(&total)
	return businessCard, total, nil
}

// 根据名片ID查询用户名片
func FindBusinessCard(id string) (*BusinessCard, error) {
	var businessCard BusinessCard
	err := Orm.Model(&BusinessCard{}).Where("id = ? and data_status=?", id, e.Enable).Find(&businessCard).Error
	if err != nil {
		return nil, err
	}
	return &businessCard, nil
}

// 更新名片
func UpdateBusinessCard(businessCard *BusinessCard, data interface{}) bool {
	err := Orm.Model(&businessCard).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

//删除名片
func DeleteBusinessCard(businessCard *BusinessCard) bool {
	err := Orm.Delete(&businessCard).Error
	if err != nil {
		return false
	}
	return true
}
