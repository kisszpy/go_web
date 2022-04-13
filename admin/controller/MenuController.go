package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/admin/service"
	"go_web/common"
	"go_web/global"
	"time"
)

var resourceService = new(service.ResourceService)

type MenuController struct {
}

func (MenuController) ShowMenus(ctx *gin.Context) {
	var list []model.Resource
	resourceService.GetAllResource(&list)
	// 模型映射
	var menuList []resp.Menu
	for _, item := range list {
		data := resp.Menu{}
		json.Unmarshal([]byte(item.Data), &data)
		menuList = append(menuList, resp.Menu{
			Id:         item.Id,
			Pid:        item.Pid,
			Name:       item.Name,
			Path:       data.Path,
			Component:  data.Component,
			Redirect:   data.Redirect,
			AlwaysShow: data.AlwaysShow,
			Meta: resp.Meta{
				Title: data.Meta.Title,
				Icon:  data.Meta.Icon,
			},
		})
	}
	// 构建树形
	var treeMenu []resp.Menu
	for _, item := range menuList {
		if isRoot(item) {
			treeMenu = append(treeMenu, findChildren(item, menuList))
		}
	}
	var menu resp.Menu
	menu.AlwaysShow = true
	menu.Redirect = "/permission/page"
	menu.Component = "Layout"
	menu.Path = "/permission"
	menu.Name = "permission"
	menu.Meta = resp.Meta{Title: "permission", Icon: "lock"}

	m := resp.Menu{
		Path:      "page",
		Component: "permission/page",
		Name:      "PagePermission",
		Meta: resp.Meta{
			Title: "PagePermission",
			Icon:  "lock",
		},
	}
	menu.Children = append(menu.Children, m)

	var menu2 resp.Menu
	menu2.AlwaysShow = true
	menu2.Redirect = "/system/user"
	menu2.Component = "Layout"
	menu2.Path = "/system"
	menu2.Name = "System"
	menu2.Meta = resp.Meta{Title: "系统管理", Icon: "lock"}

	m2 := resp.Menu{
		Path:      "user",
		Component: "system/user",
		Name:      "User",
		Meta: resp.Meta{
			Title: "用户管理",
			Icon:  "lock",
		},
	}

	m3 := resp.Menu{
		Path:      "role",
		Component: "system/role",
		Name:      "Role",
		Meta: resp.Meta{
			Title: "角色管理",
			Icon:  "lock",
		},
	}

	menu2.Children = append(menu2.Children, m2, m3)

	var menus []resp.Menu
	menus = append(menus, menu, menu2)
	common.Success(treeMenu, ctx)

}

func (MenuController) Load(context *gin.Context) {
	req := &req.IdReq{}
	context.BindJSON(req)
	result := resourceService.LoadResourceById(req.Id)
	common.Success(result, context)
}

func (c MenuController) Test(ctx *gin.Context) {
	token := global.GetToken("200")
	common.Success(token, ctx)
}

func (c MenuController) Test2(ctx *gin.Context) {
	//token := ctx.GetHeader("Auth-Token")
	//verifyResult, _ := global.Verify(token)
	//common.Success(verifyResult, ctx)
	eq := global.Toolkit.Eq("2", "2")
	fmt.Printf("%v \n", eq)
	myTime := resp.MyTime{
		CreateTime: common.Time(time.Now()),
	}
	common.Success(myTime, ctx)
}

func findChildren(root resp.Menu, menus []resp.Menu) resp.Menu {
	for _, menu := range menus {
		if root.Id == menu.Pid {
			root.Children = append(root.Children, menu)
		}
	}
	return root
}

func isRoot(item resp.Menu) bool {
	return item.Pid == 0
}
