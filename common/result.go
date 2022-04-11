package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SuccessCode = 200
const FailCode = 500

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(SuccessCode, success(data))
}
func Ok(ctx *gin.Context) {
	ctx.JSON(SuccessCode, success(make(map[string]string)))
}

func success(data interface{}) Result {
	return Result{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: data,
	}
}

func Fail(error string) Result {
	return Result{
		Code: 500,
		Msg:  "fail",
		Data: nil,
	}
}
