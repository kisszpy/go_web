package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	GDB *gorm.DB
)

func init() {
	dsn := "root:Metro#79@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("error ......... can't open db")
		return
	}
	GDB = database

}
