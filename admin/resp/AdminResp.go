package resp

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type Menu struct {
	Id         int    `json:"id"`
	Pid        int    `json:"pid"`
	AlwaysShow bool   `json:"alwaysShow"`
	Children   []Menu `json:"children"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Meta       Meta   `json:"meta"`
}
