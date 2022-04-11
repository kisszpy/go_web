package service

import (
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/dao"
	"log"
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
	dao.Insert(&user)
}

func (UserService) List(userList *[]model.User) {
	dao.FindAll(userList)
}

func (UserService) Delete(id int) {
	log.Printf("%v           \n", id)
	var user model.User
	dao.FindById(&user, id)
	dao.Delete(&user)
}
