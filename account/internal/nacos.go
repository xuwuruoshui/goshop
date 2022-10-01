package internal

type Nacos struct {
	Host string `mapstructure:"host"`
	Port uint64 `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	DataId string `mapstructure:"dataId"`
	Group string `mapstructure:"group"`
}
