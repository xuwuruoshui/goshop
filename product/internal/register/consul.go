package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"product/internal"
)

type IRegister interface {

	Register(name,id,port,int,tags []string) error
	DeRegister(serviceId string)error

}

type ConsulRegistry struct {
	Host string
	Port int
}

func NewConsulRegistry(host string,port int)ConsulRegistry{
	return ConsulRegistry{
		Host: host,
		Port: port,
	}
}

func (cr ConsulRegistry)Register(name,id string,port int,tags []string)error{
	config := api.DefaultConfig()
	consulHost := internal.AppConf.Consul.Host
	consulPort := internal.AppConf.Consul.Port
	config.Address =fmt.Sprintf("%s:%d",consulHost,consulPort)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = config.Address
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags
	serverAddr := fmt.Sprintf("http://%s:%d/health",internal.AppConf.ProductWeb.Host,internal.AppConf.ProductWeb.Port)
	check := api.AgentServiceCheck{
		HTTP:                           serverAddr,
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "20s",
	}
	agentServiceRegistration.Check = &check
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