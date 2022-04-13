package resp

type UserProfileResp struct {
	Roles        []string `json:"roles"`        // 角色
	Introduction string   `json:"introduction"` // 介绍
	Avatar       string   `json:"avatar"`       // 头像
	Name         string   `json:"name"`         // 名称
}
