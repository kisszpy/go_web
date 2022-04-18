package service

import (
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/global"
	"time"
)

type RoleResourceService struct {
}

func (RoleResourceService) SubmitPermission(req *req.SubmitPermissionReq) {
	for _, resId := range req.Resources {
		rp := &model.RolePermission{
			RoleId:     req.RoleId,
			ResourceId: resId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		var dest model.RolePermission
		global.GDB.Model(&model.RolePermission{}).Where("role_id = ? and resource_id = ?", rp.RoleId, rp.ResourceId).Find(&dest)
		if dest.Id == 0 {
			global.GDB.Model(&model.RolePermission{}).Create(rp)
		}
	}
}
