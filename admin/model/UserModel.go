package model

import (
	"time"
)

// User /**
type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Nickname   string    `json:"nickname"`
	Email      string    `json:"email"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`
}

// User tableName /**

func (User) TableName() string {
	return "t_user"
}
