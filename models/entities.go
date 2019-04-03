package models

import "time"

type User struct {
	Id         string    `json:"id"` // id
	NikeName   string    `json:"nike_name"` // 昵称
	AvatarUrl  string    `json:"avatar_url"` // 头像链接
	Email      string    `json:"email"` // 邮箱
	Phone      string    `json:"phone"` // 手机号
	Sex        int8      `json:"sex"` // 性别
	Birthday   time.Time `json:"birthday"` // 生日
	Address    string    `json:"address"` // 地址
	CreateTime time.Time `json:"create_time"` // 创建日期
	UpdateTime time.Time `json:"update_time"` // 更新日期
	DataStatus int8      `json:"data_status"` // 用户状态
}

type UserLoginMethod struct {
	Id             string    `json:"id"` // id
	UserId         string    `json:"user_id"` //  用户ID
	LoginType      string    `json:"login_type"` //  登陆方法
	Identification string    `json:"identification"` //  登陆标识
	AccessCode     string    `json:"access_code"` //  登陆密码或授权码
	Status         int8      `json:"status"` // 该登陆方式是否禁用0禁用，1开启
	CreateTime     time.Time `json:"create_time"` // 创建日期
	UpdateTime     time.Time `json:"update_time"` // 更新日期
}
