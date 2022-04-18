package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/service"
	"go_web/common"
	"go_web/global"
	"reflect"
	"strconv"
)

type RoleController struct {
}

var (
	roleService    service.RoleService
	rolePermission service.RoleResourceService
)

// Create 添加角色/**
func (RoleController) Create(ctx *gin.Context) {
	request := &req.CreateRoleReq{}
	ctx.BindJSON(request)
	roleService.Create(request)
	common.Ok(ctx)
}

// Modify 更新角色
func (RoleController) Modify(ctx *gin.Context) {
	req := &req.ModifyRoleReq{}
	ctx.BindJSON(req)
	roleService.Modify(req)
	common.Ok(ctx)
}

// List 角色列表
func (RoleController) List(ctx *gin.Context) {
	roleListReq := &req.RoleListReq{}
	ctx.BindJSON(roleListReq)
	list, _ := roleService.List(roleListReq)
	common.Success(list, ctx)
}

func (RoleController) Delete(ctx *gin.Context) {
	idReq := &req.IdReq{}
	ctx.BindJSON(idReq)
	roleService.Delete(idReq.Id)
	common.Ok(ctx)
}
func (RoleController) GetRoleMenus(ctx *gin.Context) {
	value, exists := ctx.Get(global.CONF.Jwt.Context)
	if exists {
		userId := reflect.ValueOf(value).String()
		fmt.Printf("get context value is %v \n", userId)
		uid, _ := strconv.Atoi(userId)
		menus := roleService.RoleMenus(uid)
		common.Success(menus, ctx)
	} else {
		common.Fail("未知错误", ctx)
	}
}

func (RoleController) SubmitPermission(context *gin.Context) {
	req := &req.SubmitPermissionReq{}
	context.BindJSON(req)
	rolePermission.SubmitPermission(req)
	common.Ok(context)
}
