package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/common"
	"go_web/global"
)

type LoginController struct {
}

func (LoginController) Login(ctx *gin.Context) {
	//data := "admin-token"
	loginReq := &req.LoginReq{}
	ctx.BindJSON(loginReq)
	e := loginReq.CheckArguments()
	if e != nil {
		common.Fail(e.Error(), ctx)
	}
	loginUser, err := userService.Login(loginReq)
	if err == nil {
		token := global.GetToken(string(loginUser.Id), "www.yukens.com")
		loginResponse := resp.LoginResp{
			AccessToken:  token,
			RefreshToken: token,
		}
		common.Success(loginResponse, ctx)
	} else {
		common.Fail(err.Error(), ctx)
	}

}

func (LoginController) LogOut(ctx *gin.Context) {
	common.Ok(ctx)
}
