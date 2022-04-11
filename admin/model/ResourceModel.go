package model

import "time"

type Resource struct {
	Id         uint
	Name       string
	Pid        uint
	Data       string
	CreateTime time.Time
	UpdateTime time.Time
	status     uint
}

func (Resource) TableName() string {
	return "t_resource"
}
