package router

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/controller"
)

var (
	userController  = new(controller.UserController)
	loginController = new(controller.LoginController)
	menuController  = new(controller.MenuController)
	roleController  = new(controller.RoleController)
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
	group := e.Group("/api/v1/admin")
	{
		/**
		后台登录接口
		*/
		group.POST("/login", loginController.Login)
		/**

		 */
		group.POST("/logout", loginController.LogOut)

		/**
		后台用户信息
		*/
		group.POST("/loginInfo", userController.GetUserInfo)
		/**
		系统菜单
		*/
		group.GET("/system/menus", menuController.ShowMenus)
		group.POST("/system/menus/create", menuController.Create)
		group.POST("/system/menus/modify", menuController.Modify)

		/**
		后台用户列表
		*/
		group.POST("/system/user/list", userController.List)

		/**
		添加用户
		*/
		group.POST("/system/user/create", userController.Create)

		/**
		用户删除
		*/
		group.POST("/system/user/delete", userController.Delete)
		/**
		用户删除
		*/
		group.POST("/system/user/modify", userController.Modify)
		/**
		设置角色
		*/
		group.POST("/system/user/settingRole", userController.SettingRole)
		/**
		角色部分 start ------------------------------
		*/
		group.POST("/system/role/create", roleController.Create)
		group.POST("/system/role/list", roleController.List)
		group.POST("/system/role/modify", roleController.Modify)
		group.POST("/system/role/delete", roleController.Delete)
		/**
		角色部分 end   ------------------------------
		*/
		/**
		测试框架
		*/
		group.GET("/system/user/test", userController.List)
		group.POST("/system/resource/load", menuController.Load)
		group.POST("/system/test", menuController.Test)
		group.POST("/system/test2", menuController.Test2)
	}

}
