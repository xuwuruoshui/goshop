package internal


type AccountSrv struct {
	SrvName string `mapstructure:"srvName" yaml:"srvName"`
	Host string `mapstructure:"host" yaml:"host"`
	Port int `mapstructure:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" yaml:"tags"`
}

type AccountWeb struct {
	SrvName string `mapstructure:"srvName" yaml:"srvName"`
	Host string `mapstructure:"host" yaml:"host"`
	Port int `mapstructure:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" yaml:"tags"`
}

