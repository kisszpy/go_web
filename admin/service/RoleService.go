package service

import (
	"github.com/youthlin/stream"
	"github.com/youthlin/stream/types"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/common"
	"go_web/global"
	"time"
)

type RoleService struct {
}

func (RoleService) Create(createRoleReq *req.CreateRoleReq) {
	role := model.Role{
		RoleName:   createRoleReq.RoleName,
		RoleDesc:   createRoleReq.RoleDesc,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Status:     0,
	}
	global.GDB.Model(&model.Role{}).Create(&role)
}

func (RoleService) List(query *req.RoleListReq) (common.PageResult, error) {
	pageResult := common.PageResult{}
	pageResult.Page = query.Page
	pageResult.PageSize = query.PageSize
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	// 创建db
	db := global.GDB.Model(&model.Role{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if query.RoleName != "" {
		db = db.Where("`role_name` LIKE ?", "%"+query.RoleName+"%")
	}
	db.Count(&pageResult.Total)
	var list []model.Role
	err := db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error
	respList := stream.OfSlice(list).Map(func(e types.T) types.R {
		role := e.(model.Role)
		resp := resp.RoleListResp{
			Id:         role.Id,
			RoleName:   role.RoleName,
			RoleDesc:   role.RoleDesc,
			CreateTime: role.CreateTime,
			Status:     role.Status,
		}
		return resp
	}).ToSlice()
	pageResult.List = respList
	return pageResult, err

}

func (RoleService) Delete(id int) {
	global.GDB.Model(&model.User{}).Delete("id = ?", id)
}
func (RoleService) Modify() {
	dest := &model.User{}
	global.GDB.Model(&model.User{}).Where("id = ?", 1).Find(dest)
	dest.Email = "guaga@163.com"
	global.GDB.Updates(dest)
}
