package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/common"
)

type LoginController struct {
}

func (LoginController) Login(ctx *gin.Context) {
	data := "admin-token"
	common.Success(data, ctx)
}
