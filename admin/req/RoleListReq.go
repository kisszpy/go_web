package req

import "go_web/common"

type RoleListReq struct {
	common.PageResult
	RoleName string `json:"roleName"`
}
