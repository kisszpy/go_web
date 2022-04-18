package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/admin/req"
	"go_web/admin/resp"
	"go_web/admin/service"
	"go_web/common"
	"go_web/global"
	"io/ioutil"
	"net/http"
	"strconv"
)

var loginService = new(service.LoginService)

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

func (LoginController) Test(ctx *gin.Context) {
	req := &req.MongoOptReq{}
	ctx.BindJSON(req)
	loginService.MongoTest(req)
	common.Success("ok,我是一个test服务", ctx)
}

func (LoginController) TestInvoke(context *gin.Context) {
	fmt.Printf("method is %v", "hello...............")
	for i := 0; i < 1000; i++ {
		lb := global.NacosClient.HttpPaths("tf-hades-service")
		url := fmt.Sprintf("%v%v%v", lb, "/tf-hades", "/testByGo")
		response, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(response.Body)
		common.Success(string(bytes), context)
	}

}
