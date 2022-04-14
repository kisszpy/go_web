package req

type CreateResourceReq struct {
	Pid        uint   `json:"pid"`
	AlwaysShow bool   `json:"alwaysShow"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
}