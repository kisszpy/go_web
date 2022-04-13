package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	GDB     *gorm.DB
	Toolkit *Utils
)

const (
	AuthToken = "Auth-Token"
	UserId    = "userId"
	Issuer    = "www.yukens.com"
)

func init() {
	dsn := "root:Metro#79@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("can't open db %v \n", err)
		return
	}
	GDB = database
	Toolkit = new(Utils)
}
