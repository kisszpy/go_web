package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/admin/service"
	"go_web/common"
	"go_web/global"
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
	value, exists := ctx.Get(global.UserId)
	if exists {
		fmt.Printf("current user is %v \n", value)
	}
	info := userService.GetUserInfo(value)
	roles := append([]string{}, "admin")
	profileResp := resp.UserProfileResp{
		Name:         info.Username,
		Avatar:       info.Avatar,
		Introduction: info.Introduction,
		Roles:        roles,
	}
	common.Success(profileResp, ctx)
}

func (UserController) SettingRole(ctx *gin.Context) {
	settingRoleReq := &req.SettingRoleReq{}
	ctx.BindJSON(settingRoleReq)
	err := userService.SettingRole(settingRoleReq)
	if err != nil {
		common.Fail(err.Error(), ctx)
	} else {
		common.Ok(ctx)
	}
}
