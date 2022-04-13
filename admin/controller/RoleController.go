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

func (RoleController) List(ctx *gin.Context) {

}

func (RoleController) Delete(ctx *gin.Context) {

}
