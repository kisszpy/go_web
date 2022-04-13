package service

import (
	"errors"
	"github.com/youthlin/stream"
	"github.com/youthlin/stream/types"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/common"
	"go_web/global"
	"reflect"
	"time"
)

type UserService struct {
}

func (UserService) Create(createReq *req.UserCreateReq) {
	if createReq.Password == createReq.ConfirmPassword {
		println("==================")
	}
	user := model.User{
		Username:   createReq.Username,
		Password:   createReq.Password,
		Email:      createReq.Email,
		Nickname:   createReq.Nickname,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Status:     0,
	}
	global.GDB.Model(&model.User{}).Create(&user)
}

func (UserService) List(query req.QueryUserReq) (common.PageResult, error) {
	pageResult := common.PageResult{}
	pageResult.Page = query.Page
	pageResult.PageSize = query.PageSize
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	// 创建db
	db := global.GDB.Model(&model.User{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if query.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+query.Username+"%")
	}
	db.Count(&pageResult.Total)
	var list []model.User
	err := db.Limit(limit).Offset(offset).Order("id desc").Find(&list).Error
	respList := stream.OfSlice(list).Map(func(e types.T) types.R {
		user := e.(model.User)
		resp := resp.UserListResp{
			Id:         user.Id,
			Nickname:   user.Nickname,
			Username:   user.Username,
			Email:      user.Email,
			CreateTime: user.CreateTime,
			Status:     user.Status,
		}
		return resp
	}).ToSlice()
	pageResult.List = respList
	return pageResult, err

}

func (UserService) Delete(id int) {
	global.GDB.Model(&model.User{}).Delete("id = ?", id)
}
func (UserService) Modify() {
	dest := &model.User{}
	global.GDB.Model(&model.User{}).Where("id = ?", 1).Find(dest)
	dest.Email = "guaga@163.com"
	global.GDB.Updates(dest)
}

func (UserService) Login(loginReq *req.LoginReq) (*model.User, error) {
	dest := &model.User{}
	global.GDB.Model(&model.User{}).Where("username = ?", loginReq.Username).Find(dest)
	if reflect.DeepEqual(dest, &model.User{}) {
		return nil, errors.New("用户不存在")
	} else if dest.Password == loginReq.Password {
		return dest, nil
	} else {
		return nil, errors.New("用户名或密码错误")
	}
}

//SettingRole 设定角色 /*
func (UserService) SettingRole(roleReq *req.SettingRoleReq) error {
	dest := &model.UserRole{}
	global.GDB.Model(&model.UserRole{}).Where("role_id = ? and user_id = ?", roleReq.RoleId, roleReq.UserId).First(dest)
	if dest != nil {
		return errors.New("角色已经存在，不能重复设定")
	}
	global.GDB.Model(&model.UserRole{}).Create(&model.UserRole{
		UserId:     roleReq.UserId,
		RoleId:     roleReq.RoleId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Status:     0,
	})
	return nil
}
