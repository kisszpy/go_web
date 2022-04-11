package main

import (
	"github.com/gin-gonic/gin"
	"go_web/router"
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
	router.UserRouter(e)
	router.SiteRouter(e)
	router.AdminRouter(e)
	e.Run("localhost:9999")
}
