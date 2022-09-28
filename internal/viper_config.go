package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

var ViperConf ViperConfig
var fileName = "./dev-config.yaml"

func init(){
	v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	v.Unmarshal(&ViperConf)
	fmt.Println(ViperConf)
	fmt.Println("Viper初始化完成")
	InitRedis()

}

type ViperConfig struct {
	Redis Redis `mapstructure:"redis"`
	Consul Consul `mapstructure:"consul"`
	AccountSrv AccountSrv `mapstructure:"accountSrv"`
	AccountWeb AccountWeb `mapstructure:"accountWeb"`
}
