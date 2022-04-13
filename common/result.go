package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, success(data))
}
func Ok(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, success(make(map[string]string)))
}

func success(data interface{}) Result {
	return Result{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: data,
	}
}

func Fail(error string, ctx *gin.Context) {
	fail := Result{
		Code: 500,
		Msg:  error,
		Data: make(map[string]string),
	}
	ctx.JSON(500, fail)
}
