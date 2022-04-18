package global

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	MongoClient *mongo.Client
	RedisClient *redis.Conn
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
	// use config file to load mongodb
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user:Test#1234@172.20.30.23:8635,172.20.30.81:8635/seagull?authSource=admin&replicaSet=replica"))
	client, err := mongo.NewClient(options.Client().ApplyURI(CONF.MongoDb.Uri))
	if err == nil {
		e := client.Connect(context.TODO())
		if e == nil {
			client.Ping(context.TODO(), nil)
		}
		MongoClient = client
		defer client.Disconnect(context.TODO())
	} else {
		// fmt.Printf("error to connect mongodb %v \n", err)
		panic("error to load mongodb")
	}
}

func initRedis() {
	connStr := fmt.Sprintf("%v:%v", CONF.Redis.Host, CONF.Redis.Port)
	client, err := redis.Dial("tcp", connStr, redis.DialPassword(CONF.Redis.Pass))
	if err == nil {
		RedisClient = &client
		client.Do("set", "laohe", "唐荣的电脑是windows64")
	} else {
		fmt.Printf("%v\n", err)
		panic("redis connect error ....")
	}
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
