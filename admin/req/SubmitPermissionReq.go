package req

type SubmitPermissionReq struct {
	RoleId    int   `json:"roleId"`
	Resources []int `json:"resources"`
}
