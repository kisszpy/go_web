package router

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/controller"
)

var (
	userController  = new(controller.UserController)
	loginController = new(controller.LoginController)
	menuController  = new(controller.MenuController)
)

func SiteRouter(e *gin.Engine) {
	//e.GET("/index", controller.Index)
}

func UserRouter(e *gin.Engine) {
	/**
	用户登录接口
	*/
	//e.GET("/api/v1/login", controller.Login)

	/**
	用户注册接口
	*/
	//e.POST("/api/v1/register", controller.Register)
}
func AdminRouter(e *gin.Engine) {
	/**
	后台登录接口
	*/
	e.POST("/api/v1/admin/login", loginController.Login)

	/**
	后台用户信息
	*/
	e.GET("/api/v1/admin/loginInfo", userController.GetUserInfo)
	/**
	系统菜单
	*/
	e.GET("/api/v1/admin/system/menus", menuController.ShowMenus)

	/**
	后台用户列表
	*/
	e.POST("/api/v1/admin/system/user/list", userController.List)

	/**
	添加用户
	*/
	e.POST("/api/v1/admin/system/user/create", userController.Create)

	/**
	用户删除
	*/
	e.POST("/api/v1/admin/system/user/delete", userController.Delete)

	/**
	测试框架
	*/
	e.GET("/api/v1/admin/system/user/test", userController.List)
}
