package model

import "time"

type Resource struct {
	Id         int
	Name       string
	Pid        int
	Data       string
	CreateTime time.Time
	UpdateTime time.Time
	status     uint
}

func (Resource) TableName() string {
	return "t_resource"
}
