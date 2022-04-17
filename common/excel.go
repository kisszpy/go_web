package common

import (
	"github.com/gin-gonic/gin"
)

type Excel struct {
}

func (Excel) ImportExcel(ctx *gin.Context) {
	header, err := ctx.FormFile("file")
	if err != nil {
		Fail("导入文件失败", ctx)
	}
	ctx.SaveUploadedFile(header, header.Filename)
	Success("导入文件成功", ctx)
}
