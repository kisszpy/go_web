package controller

import (
	"github.com/gin-gonic/gin"
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
	req := req.QueryUserReq{}
	ctx.BindJSON(&req)
	list, _ := userService.List(req)
	common.Success(list, ctx)
}
func (UserController) Modify(ctx *gin.Context) {
	userService.Modify()
	common.Ok(ctx)
}

func (UserController) GetUserInfo(ctx *gin.Context) {
	ctx.GetHeader("")
	roles := append([]string{}, "admin")
	profileResp := resp.UserProfileResp{
		Name:         "Super Admin",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        roles,
	}
	common.Success(profileResp, ctx)
}
