package model

import "time"

type UserRole struct {
	Id         int       `json:"id"`
	RoleId     int       `json:"roleId"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}

func (UserRole) TableName() string {
	return "t_user_role"
}
