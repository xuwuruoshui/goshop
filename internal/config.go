package internal


type AccountSrv struct {
	SrvName string `mapstructure:"srvName" json:"srvName"`
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}

type AccountWeb struct {
	SrvName string `mapstructure:"srvName" json:"srvName"`
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}

