package global

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"strconv"
)

var INamingClient naming_client.INamingClient

type Nacos struct {
}

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (Nacos) InitNacos() {
	config := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/log",
		CacheDir:            "/tmp/cache",
		LogLevel:            "debug",
	}
	serverConfig := []constant.ServerConfig{
		{
			Scheme:      "http",
			ContextPath: "/nacos",
			IpAddr:      "10.6.2.180",
			Port:        8848,
		},
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &config,
			ServerConfigs: serverConfig,
		})
	if err != nil {
		fmt.Printf("init error %v \n", err)
	}
	INamingClient = namingClient
}

func (Nacos) Register() {
	INamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          LocalIP(),
		Port:        9999,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: "DEFAULT",
		ServiceName: "go-web-service",
		GroupName:   "DEFAULT_GROUP",
		Ephemeral:   true,
	})
}

func (Nacos) HttpPaths(serverName string) (httpPath string) {
	instance, _ := INamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serverName,
	})
	httpPath = "http://" + instance.Ip + ":" + strconv.Itoa(int(instance.Port))
	return
}
