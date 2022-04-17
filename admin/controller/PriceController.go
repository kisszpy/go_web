package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/model"
	"go_web/admin/service"
	"go_web/common"
	"go_web/global"
)

type PriceController struct {
}

var priceService = new(service.PriceService)
var testService = new(service.TestService)

func (PriceController) Test(*gin.Context) {
	go testService.Frozen(20)
}

func (PriceController) Import(ctx *gin.Context) {
	global.Excel.ImportExcel(ctx)
}

func (PriceController) Create(ctx *gin.Context) {
	priceModel := model.PriceModel{
		Sku:     "1002",
		SkuName: "牛肉",
	}
	priceService.InsertPrice(&priceModel)
	common.Ok(ctx)
}
