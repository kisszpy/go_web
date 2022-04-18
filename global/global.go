package global

import (
	"fmt"
	"github.com/spf13/viper"
	"go_web/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// 系统全局变量
var (
	GDB         *gorm.DB
	Toolkit     Utils
	Excel       common.Excel
	CONF        *AppConfig
	NacosClient *Nacos
)

// 系统常量
const (
	AuthToken  = "Auth-Token"
	UserId     = "userId"
	ConfigFile = "config.yaml"
	ConfigType = "yaml"
)

func init() {
	initConfig()      // 初始化配置文件
	initNacosServer() // 注册到nacos
	initMysql()       // 初始化mysql
	initMongoDb()     // 初始化mongodb
	initRedis()       // 初始化redis

}

func initNacosServer() {
	var nacosInstance Nacos
	nacosInstance.InitNacos()
	nacosInstance.Register()
	NacosClient = &nacosInstance
}

func initMongoDb() {

}

func initRedis() {

}

func initMysql() {
	// use config file to load dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/go_db?charset=utf8mb4&parseTime=true",
		CONF.Mysql.Username,
		CONF.Mysql.Password,
		CONF.Mysql.Host,
		strconv.Itoa(CONF.Mysql.Port))
	// dsn := "root:Metro#79@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("can't open db %v \n", err)
		return
	}
	GDB = database
}

// bootstrap
func initConfig() {
	v := viper.New()
	v.SetConfigType(ConfigType)
	v.SetConfigFile(ConfigFile)
	err := v.ReadInConfig()
	if err == nil {
		config := AppConfig{}
		v.Unmarshal(&config)
		CONF = &config
	}

}
