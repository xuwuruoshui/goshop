package internal


type StockSrv struct {
	SrvName string `mapstructure:"srvName" yaml:"srvName"`
	Host string `mapstructure:"host" yaml:"host"`
	Port int `mapstructure:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" yaml:"tags"`
}

type StockWeb struct {
	SrvName string `mapstructure:"srvName" yaml:"srvName"`
	Host string `mapstructure:"host" yaml:"host"`
	Port int `mapstructure:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" yaml:"tags"`
}

