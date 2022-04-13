package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/service"
	"go_web/common"
)

type RoleController struct {
}

var (
	roleService service.RoleService
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

}

// List 角色列表
func (RoleController) List(ctx *gin.Context) {
	roleListReq := &req.RoleListReq{}
	ctx.BindJSON(roleListReq)
	list, _ := roleService.List(roleListReq)
	common.Success(list, ctx)
}

func (RoleController) Delete(ctx *gin.Context) {

}
