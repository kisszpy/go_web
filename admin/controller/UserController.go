package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/admin/service"
	"go_web/common"
)

var userService = new(service.UserService)

type UserController struct {
}

func (UserController) Create(ctx *gin.Context) {
	createReq := req.UserCreateReq{}
	ctx.BindJSON(&createReq)
	userService.Create(&createReq)
	common.Success("ok", ctx)
}
func (UserController) Delete(ctx *gin.Context) {
	var idReq req.IdReq
	ctx.BindJSON(&idReq)
	userService.Delete(idReq.Id)
	common.Ok(ctx)
}

func (UserController) List(ctx *gin.Context) {
	var userList []model.User
	var result []resp.UserListResp
	userService.List(&userList)
	for _, x := range userList {
		f := func(x model.User) resp.UserListResp {
			var u resp.UserListResp
			u.Username = x.Username
			u.Nickname = x.Nickname
			u.CreateTime = x.CreateTime
			u.Email = x.Email
			u.Id = x.Id
			u.Status = x.Status
			return u
		}
		result = append(result, f(x))
	}
	common.Success(result, ctx)
}
func (UserController) Modify(ctx *gin.Context) {

}

func (UserController) GetUserInfo(ctx *gin.Context) {

	/**
	'admin-token': {
	    roles: ['admin'],
	    introduction: 'I am a super administrator',
	    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
	    name: 'Super Admin'
	  },
	*/
	var strings []string
	strings = append(strings, "admin")
	var data resp.LoginResp
	data.Roles = strings
	data.Introduction = "I am a super administrator"
	data.Avatar = "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F48%2F2c%2F90%2F482c90a6491541e8ceda713b31642a5c.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1652187967&t=9275c22a07dc968dc322b9cfe87c1683"
	data.Name = "Super Admin"
	common.Success(data, ctx)
}
