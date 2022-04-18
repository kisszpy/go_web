package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go_web/admin/model"
	"go_web/admin/req"
	"go_web/global"
)

type LoginService struct {
}

func (LoginService) Login(ctx *gin.Context) {
	println("================= login test")
}

func (LoginService) MongoTest(req *req.MongoOptReq) {
	if req.Name == "add" {
		global.MongoClient.Database("seagull").Collection("oms").InsertOne(context.TODO(), &model.Student{
			StuNo:     "002",
			Name:      "傻逼小司",
			ClassName: "弱智一班",
			Grade:     "幼儿园",
			Score:     0,
		})
	} else if req.Name == "modify" {
		// 修改
		global.MongoClient.Database("seagull").Collection("oms").UpdateOne(context.TODO(), bson.M{
			"name": "傻逼小司",
		}, &bson.M{"score": 1})
	} else if req.Name == "delete" {
		// 删除
		global.MongoClient.Database("seagull").Collection("oms").DeleteOne(context.TODO(), bson.M{
			"name": "google",
		})
	} else if req.Name == "find" {
		result, _ := global.MongoClient.Database("seagull").Collection("oms").Find(context.TODO(), bson.M{})
		data := &[]interface{}{}
		result.All(context.TODO(), data)
		fmt.Printf("result is %v", data)
	}
}
