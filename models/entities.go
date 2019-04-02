package models

import "time"

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	AvatarUrl  string    `json:"avatar_url"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Sex        int8      `json:"sex"`
	Birthday   time.Time `json:"birthday"`
	Address    string    `json:"address"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DataStatus int8      `json:"data_status"`
}
