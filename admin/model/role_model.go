package model

import "time"

type Role struct {
	Id         int       `json:"id"`
	RoleName   string    `json:"roleName"`
	RoleDesc   string    `json:"roleDesc"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}

func (Role) TableName() string {
	return "t_role"
}
