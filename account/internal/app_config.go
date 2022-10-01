package internal

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var NacosConf *Nacos = &Nacos{}
var AppConf *App = &App{}
var fileName = "./dev-config.yml"

func initNacos(){
	v := viper.New()
	v.SetConfigFile(fileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(NacosConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(NacosConf)
	fmt.Println("Viper初始化Nacos完成")
}

func initAppConf(){

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:NacosConf.Host,
			Port: NacosConf.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConf.NameSpace,
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
		DataId: NacosConf.DataId,
		Group:  NacosConf.Group,
	})
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(config), &AppConf)
	if err != nil {
		panic(err)
	}
}

func init(){

	// 1.初始化 Nacos
	initNacos()
	// 2.初始化 App配置
	initAppConf()
	InitRedis()

}

type App struct {
	DataBase DataBase `mapstructure:"database" yaml:"database"`
	Redis Redis `mapstructure:"redis" yaml:"redis"`
	Consul Consul `mapstructure:"consul" yaml:"consul"`
	AccountSrv AccountSrv `mapstructure:"accountSrv" yaml:"accountSrv"`
	AccountWeb AccountWeb `mapstructure:"accountWeb" yaml:"accountWeb"`
	JWT JWT `mapstructure:"jwt" yaml:"jwt"`
}
