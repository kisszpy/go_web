package model

import (
	"time"
)

type UserInfo struct {
	Id         uint
	Username   string
	Password   string
	Email      string
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     uint
}

func (UserInfo) TableName() string {
	return "user_info_v2"
}
