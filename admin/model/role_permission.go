package model

import "time"

type RolePermission struct {
	Id         int
	RoleId     int
	ResourceId int
	CreateTime time.Time
	UpdateTime time.Time
	Status     int
}
