package resp

type LoginResp struct {
	/**
	'admin-token': {
	    roles: ['admin'],
	    introduction: 'I am a super administrator',
	    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
	    name: 'Super Admin'
	  },
	*/
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type Menu struct {
	Id         uint   `json:"id"`
	Pid        uint   `json:"pid"`
	AlwaysShow bool   `json:"alwaysShow"`
	Children   []Menu `json:"children"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Meta       Meta   `json:"meta"`
}

type User struct {
	Id         int    `json:"id"`
	Avatar     string `json:"avatar"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Dept       string `json:"dept"`
	CreateTime string `json:"createTime"`
}
