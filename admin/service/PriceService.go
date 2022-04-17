package service

import (
	"go_web/admin/model"
	"go_web/global"
	"log"
)

type PriceService struct {
}

func (PriceService) InsertPrice(dest *model.PriceModel) {
	log.Printf("dest value is %v", dest)
	global.GDB.Model(&model.PriceModel{}).Create(dest)
}
