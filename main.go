package main

import (
	"github.com/gin-gonic/gin"
	"go_web/controller"
	"go_web/router"
	"log"
	"net/http"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/**")
	e.Static("/www", "./www")
	router.UserRouter(e)
	router.SiteRouter(e)
	e.GET("/hello", controller.Hello)
	e.GET("/login", Login)
	e.GET("/register", Register)
	e.GET("/oop", controller.OOP)
	e.POST("/doRegister", DoRegister)
	e.Run()
}

func DoRegister(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	log.Printf("%v---------%v", username, password)
	context.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
		"password": password,
		"list":     []string{"a", "b"},
	})
}

func Register(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", nil)
}

func Login(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func Hello(context *gin.Context) {
	method := context.Request.Method
	context.String(http.StatusOK, "helfdfds lo%s", method)
}
