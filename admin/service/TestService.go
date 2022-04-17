package service

import (
	"go_web/admin/model"
	"go_web/global"
)

type TestService struct {
}

func (TestService) Frozen(num int) {
	test := model.Test{}
	tx := global.GDB.Model(&test)
	tx.Begin()
	dest := &model.Test{}
	tx.Where("id = 1").Find(dest).UpdateColumn("frozen", dest.Frozen+num).UpdateColumn("available", dest.Available-num)
	tx.Commit()
}
