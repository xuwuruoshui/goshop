package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"stock/internal"
)

type IRegister interface {

	Register(name,id,port,int,tags []string) error
	DeRegister(serviceId string)error

}


type RPCType int
const(
	HTTP RPCType = iota
	GRPC
)

type ConsulRegistry struct {
	Host string
	Port int
	RPCType RPCType
}

func NewConsulRegistry(host string,port int,rpcType RPCType)ConsulRegistry{
	return ConsulRegistry{
		Host: host,
		Port: port,
		RPCType: rpcType,
	}
}

func (cr ConsulRegistry)Register(name,id,host string,port int,tags []string)error{
	config := api.DefaultConfig()
	consulHost := internal.AppConf.Consul.Host
	consulPort := internal.AppConf.Consul.Port
	config.Address =fmt.Sprintf("%s:%d",consulHost,consulPort)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Error(err)
		return err
	}



	var check api.AgentServiceCheck
	switch cr.RPCType {
	case GRPC:
		serverAddr := fmt.Sprintf("%s:%d",host,port)
		check = api.AgentServiceCheck{
			GRPC: serverAddr,
			Timeout: "3s",
			Interval: "1s",
			DeregisterCriticalServiceAfter: "5s",
		}
	case HTTP:
		serverAddr := fmt.Sprintf("http://%s:%d/health",host,port)
		check = api.AgentServiceCheck{
			HTTP:                           serverAddr,
			Timeout:                        "3s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "20s",
		}
	default:
		zap.S().Panic("consul心跳检查配置失败")
	}

	agentServiceRegistration := &api.AgentServiceRegistration{
		// 服务名称(可以相同，但id必须不一样)
		Name: name,
		// 每个实例的id
		ID: id,
		Port: port,
		Tags: tags,
		Address: host,
		Check: &check,
	}

	zap.S().Info("当前节点端口:",port)
	zap.S().Info("当前节点ID:",id)
	return client.Agent().ServiceRegister(agentServiceRegistration)
}


func (cr ConsulRegistry)DeRegister(serviceId string)error{
	config := api.DefaultConfig()
	consulHost := internal.AppConf.Consul.Host
	consulPort := internal.AppConf.Consul.Port
	config.Address =fmt.Sprintf("%s:%d",consulHost,consulPort)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	return client.Agent().ServiceDeregister(serviceId)
}