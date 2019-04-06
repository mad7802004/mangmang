package models

import (
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"
)

// 用户表
type User struct {
	Id         string         `json:"id"gorm:"primary_key"` // id
	Name       string         `json:"name"`                 // 昵称
	AvatarUrl  string         `json:"avatar_url"`           // 头像链接
	Email      string         `json:"email"`                // 邮箱
	Phone      string         `json:"phone"`                // 手机号
	Sex        int8           `json:"sex"gorm:"default"`    // 性别
	Birthday   utils.JSONDate `json:"birthday"`             // 生日
	Address    string         `json:"address"`              // 地址
	CreateTime utils.JSONTime `json:"-"`                    // 创建日期
	UpdateTime utils.JSONTime `json:"-"`                    // 更新日期
	DataStatus int8           `json:"-"gorm:"default" `     // 用户状态
}

// 用户登录授权表
type UserLoginMethod struct {
	Id             string         `json:"id"gorm:"primary_key"` // id
	UserId         string         `json:"user_id"`              //  用户ID
	LoginType      string         `json:"login_type"`           //  登陆方法
	Identification string         `json:"identification"`       //  登陆标识
	AccessCode     string         `json:"-"`                    //  登陆密码或授权码
	CreateTime     utils.JSONTime `json:"-"`                    // 创建日期
	UpdateTime     utils.JSONTime `json:"-"`                    // 更新日期
	DataStatus     int8           `json:"-"gorm:"default"`      // 该登陆方式是否禁用0禁用，1开启
}

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
	DataStatus int8           `json:"data_status"`          // 0删除，1有效
}

// 查询手机号是否被注册使用
func IsExistPhone(phone string) bool {

	_, err := FindPhoneLoginMethod(phone)
	if err != nil {
		return true
	}
	return false

}

// 根据手机号查询用户手机登录方式信息
func FindPhoneLoginMethod(phone string) (*UserLoginMethod, error) {
	var loginMethod UserLoginMethod
	err := Orm.Model(&UserLoginMethod{}).
		Where("Identification =? and data_status=?", phone, e.Enable).Find(&loginMethod).Error
	if err != nil {
		return nil, err
	}
	return &loginMethod, nil
}

// 根据用户ID查询手机登录方式信息
func FindUserIdLoginMethod(userId string) (*UserLoginMethod, error) {
	var loginMethod UserLoginMethod

	err := Orm.Model(&UserLoginMethod{}).
		Where("user_id =? and login_type =? and data_status=?", userId, e.LoginPhone, e.Enable).
		Find(&loginMethod).Error
	if err != nil {
		return nil, err
	}
	return &loginMethod, nil
}

// 根据ID查询用户信息
func FindUserIdInfo(userId string) (*User, error) {
	var user User
	err := Orm.Model(&User{}).Where("id = ? and data_status=?", userId, e.Enable).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(data ...interface{}) bool {
	tx := Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false
	}
	for _, model := range data {

		if err := tx.Create(model).Error; err != nil {
			tx.Rollback()
			return false
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return false
	}
	return true

}

// 更新用户信息
func UpdateUserInfo(user *User, data interface{}) bool {
	err := Orm.Model(&user).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 更新用户手机登录密码
func UpdateUserPassWord(loginMethod *UserLoginMethod, passWord string) bool {
	err := Orm.Model(&loginMethod).
		Updates(map[string]interface{}{"access_code": utils.Md5Encrypt(passWord)}).Error
	if err != nil {
		return false
	}
	return true

}
