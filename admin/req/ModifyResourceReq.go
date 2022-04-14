package req

type ModifyResourceReq struct {
	Id         int    `json:"id"`
	Pid        int    `json:"pid"`
	AlwaysShow bool   `json:"alwaysShow"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
}
