package service

import (
	"encoding/json"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/global"
	"time"
)

type ResourceService struct {
}

func (ResourceService) GetAllResource(list *[]model.Resource) {
	global.GDB.Model(&model.Resource{}).Find(&list)
}

func (ResourceService) LoadResourceById(id int) model.Resource {
	resource := model.Resource{}
	global.GDB.Model(&model.Resource{}).Where("id = ?", id).Find(&resource)
	return resource
}

func (ResourceService) Create(req *req.CreateResourceReq) {
	// {"meta": {"icon": "el-icon-setting", "title": "系统管理"}, "name": "System", "path": "/system", "redirect": "/system", "component": "Layout", "alwaysShow": true}
	global.GDB.Model(&model.Resource{}).Where("name = ?", req.Name)
	m := make(map[string]interface{})
	m["name"] = req.Name
	m["component"] = req.Component
	m["path"] = req.Path
	m["redirect"] = req.Redirect
	m["alwaysShow"] = req.AlwaysShow
	meta := make(map[string]interface{})
	meta["title"] = req.Title
	meta["icon"] = req.Icon
	m["meta"] = meta
	byteArray, err := json.Marshal(&m)
	if err == nil {
		global.GDB.Model(&model.Resource{}).Create(&model.Resource{
			Pid:        req.Pid,
			Name:       req.Name,
			Data:       string(byteArray),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}

}

func (ResourceService) Modify(req *req.ModifyResourceReq) {
	// find 修改
	m := make(map[string]interface{})
	m["name"] = req.Name
	m["component"] = req.Component
	m["path"] = req.Path
	m["redirect"] = req.Redirect
	m["alwaysShow"] = req.AlwaysShow
	meta := make(map[string]interface{})
	meta["title"] = req.Title
	meta["icon"] = req.Icon
	m["meta"] = meta
	byteArray, _ := json.Marshal(&m)
	dest := &model.Resource{}
	db := global.GDB.Model(&model.Resource{})
	db.Where("id = ?", req.Id).Find(dest)
	dest.Pid = req.Pid
	dest.Name = req.Name
	dest.Data = string(byteArray)
	db.Updates(dest)
}
