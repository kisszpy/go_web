package router

import (
	"github.com/gin-gonic/gin"
	"go_web/controller"
)

func SiteRouter(e *gin.Engine) {
	e.GET("/index", controller.Index)
}

func UserRouter(e *gin.Engine) {
	/**
	用户登录接口
	*/
	e.GET("/api/v1/login", controller.Login)

	/**
	用户注册接口
	*/
	e.POST("/api/v1/register", controller.Register)
}
