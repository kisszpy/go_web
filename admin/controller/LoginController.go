package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/common"
	"go_web/global"
	"strconv"
)

type LoginController struct {
}

func (LoginController) Login(ctx *gin.Context) {
	loginReq := &req.LoginReq{}
	ctx.BindJSON(loginReq)
	e := loginReq.CheckArguments()
	if e != nil {
		common.Fail(e.Error(), ctx)
	}
	loginUser, err := userService.Login(loginReq)
	if err == nil {
		token := global.GetToken(strconv.Itoa(loginUser.Id))
		common.Success(resp.LoginResp{
			AccessToken:  token,
			RefreshToken: token,
		}, ctx)
	} else {
		common.Fail(err.Error(), ctx)
	}

}

func (LoginController) LogOut(ctx *gin.Context) {
	common.Ok(ctx)
}

func (LoginController) Test(context *gin.Context) {
	common.Success("ok,我是一个test服务", context)
}
