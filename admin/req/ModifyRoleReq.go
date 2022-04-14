package req

type ModifyRoleReq struct {
	Id       int    `json:"id"`
	RoleName string `json:"roleName"`
	RoleDesc string `json:"roleDesc"`
}
