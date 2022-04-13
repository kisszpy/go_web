package service

import (
	"go_web/admin/model"
	"go_web/global"
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
