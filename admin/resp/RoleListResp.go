package resp

import (
	"time"
)

type RoleListResp struct {
	Id         int       `json:"id"`
	RoleName   string    `json:"roleName"`
	RoleDesc   string    `json:"roleDesc"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
}
