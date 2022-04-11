package resp

import (
	"time"
)

type UserListResp struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Email      string    `json:"email"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
}
