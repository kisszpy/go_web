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
	common.Success(treeMenu, ctx)

}

func (MenuController) Load(context *gin.Context) {
	req := &req.IdReq{}
	context.BindJSON(req)
	fmt.Printf("request id is %v", req.Id)
	result := resourceService.LoadResourceById(req.Id)
	var data = resp.LoadMenuOfDataResp{
		Id:   result.Id,
		Pid:  result.Pid,
		Name: result.Name,
		Data: result.Data,
	}
	common.Success(data, context)
}

func (MenuController) Test(ctx *gin.Context) {
	token := global.GetToken("200")
	common.Success(token, ctx)
}

func (MenuController) Test2(ctx *gin.Context) {
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

func (MenuController) Create(ctx *gin.Context) {
	req := &req.CreateResourceReq{}
	ctx.BindJSON(req)
	resourceService.Create(req)
	common.Ok(ctx)
}

func (MenuController) Modify(ctx *gin.Context) {
	resourceReq := &req.ModifyResourceReq{}
	ctx.BindJSON(resourceReq)
	resourceService.Modify(resourceReq)
	common.Ok(ctx)
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
