package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main(){


	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:"192.168.0.132",
			Port: 8848,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         "44fd93b5-beaf-43ed-a2bd-19d7e2c82c4a",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端
	// 方法1
	//configClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	// 方法2
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "account.yml",
		Group:  "dev",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
