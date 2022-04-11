package service

import (
	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (LoginService) Login(ctx *gin.Context) {
	println("================= login test")
}
