package service

import (
	"go_web/admin/model"
	"go_web/dao"
)

type ResourceService struct {
}

func (ResourceService) GetAllResource(list *[]model.Resource) {
	dao.FindAll(&list)
}
