package service

import (
	"encoding/json"
	"go_web/admin/model"
	"go_web/admin/resp"
)

type Tree struct {
}

func (Tree) BuildTree(list []model.Resource) []resp.Menu {
	// 模型映射
	var menuList []resp.Menu
	for _, item := range list {
		data := resp.Menu{}
		json.Unmarshal([]byte(item.Data), &data)
		menuList = append(menuList, resp.Menu{
			Id:         item.Id,
			Pid:        item.Pid,
			Label:      item.Name,
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
	return treeMenu
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
