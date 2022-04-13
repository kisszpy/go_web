package req

type CreateRoleReq struct {
	RoleName string `json:"roleName"`
	RoleDesc string `json:"roleDesc"`
}
