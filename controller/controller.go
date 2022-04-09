package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/common"
	"go_web/dao"
	"go_web/model"
	"log"
	"net/http"
	"time"
)

type Animal struct {
	Weight int
	Height int
	Age    int
}
type Cat struct {
	Animal
	No string
}
type Run interface {
	Exec()
}
type Dog struct {
	Animal
	Leg int
}

func (Cat) Run() {
	fmt.Printf("cat run .... \n ")
}
func (Dog) Run() {
	fmt.Printf("dog run ..... \n ")
}

func OOP(ctx *gin.Context) {
	cat := Cat{Animal{Age: 2, Height: 180, Weight: 30}, "100"}
	cat.Run()

	dog := Dog{Animal{Age: 2, Height: 200, Weight: 30}, 4}
	dog.Run()
	info := model.UserInfo{}
	dao.FindById(&info)
	info.Username = "email@999999"
	info.CreateTime = time.Now()
	info.UpdateTime = time.Now()
	dao.Update(&info)

}

func Login(ctx *gin.Context) {
	log.Printf("%v \n", ctx)
}

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
func Register(ctx *gin.Context) {
	log.Printf("register method invoke ... ")
	var userInfo model.UserInfo
	ctx.BindJSON(&userInfo)
	userInfo.CreateTime = time.Now()
	userInfo.UpdateTime = time.Now()
	userInfo.Password = common.Md5("123456")
	dao.Insert(&userInfo)
	common.Success("ok", ctx)
	// ctx.JSON(200, "ok")
}

func Hello(ctx *gin.Context) {
	m := &model.UserInfo{}
	dao.FindById(m)
	fmt.Printf("%v", m)
	ctx.JSON(http.StatusOK, gin.H{
		"data": m,
		"time": time.Now().Format("2006/01/02 15:04:05"),
	})
	//userInfo := model.UserInfo{
	//	Username:   "java",
	//	Password:   "java",
	//	Email:      "java@java.com",
	//	CreateTime: time.Now(),
	//	UpdateTime: time.Now(),
	//	Status:     0,
	//}
	// dao.Insert(&userInfo)

}
