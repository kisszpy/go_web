package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/common"
	"go_web/global"
	"go_web/router"
	"log"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Methods", "*")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/**")
	e.Static("/www", "./www")
	e.Use(Cors())
	e.Use(MiddleWare())
	// 用户路由
	router.UserRouter(e)
	// 网站路由
	router.SiteRouter(e)
	// 后台路由
	router.AdminRouter(e)
	// 价格路由
	router.PriceRouter(e)
	e.Run("localhost:9999")
}

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		uri := context.Request.RequestURI
		if !isWhiteList(uri) {
			token := context.GetHeader(global.AuthToken)
			if token == "" {
				println("token is empty")
				common.Fail("请先登录方可操作", context)
				context.AbortWithStatus(200)
				return
			}
			registeredClaims, err := global.Verify(token)
			if err == nil {
				userId := registeredClaims.Subject
				context.Set(global.UserId, userId)
				//context.Next()
			} else {
				common.Fail(err.Error(), context)
				context.AbortWithStatus(200)
				return
			}

		} else {
			// 白名单无须校验token
			context.Next()
		}
		log.Printf("request uri is %v", uri)
		log.Printf("req uri is %v", uri)
	}
}
func isWhiteList(uri string) bool {
	fmt.Printf("uri is %v \n", uri)
	whiteList := []string{
		"/api/v1/admin/login",
	}
	for _, item := range whiteList {
		if item == uri {
			return true
		}
	}
	return false
}
