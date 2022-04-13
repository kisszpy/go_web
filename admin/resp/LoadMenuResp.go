package resp

type LoadMenuResp struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Pid        int    `json:"pid"`
	Icon       string `json:"icon"`
	Title      string `json:"title"`
	Component  string `json:"component"`
	Path       string `json:"path"`
	AlwaysShow bool   `json:"always_show"`
}
