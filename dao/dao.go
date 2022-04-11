package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	dsn := "root:Metro#79@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("error ......... can't open db")
		return
	}
	db = database

}
func Insert(model interface{}) {
	db.Create(model)
}

func FindById(model interface{}, id int) {
	db.Where("id = ?", id).Find(model)
}
func Update(model interface{}) {
	db.Save(model)
}

func FindAll(list interface{}) {
	db.Order("id desc").Find(list)
}
func Delete(mode interface{}) {
	db.Delete(mode)
}
