package internal

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

// 服务注册、心跳
func Reg(host,name,id string,port int,tags []string)error{
	config := api.DefaultConfig()
	h := ViperConf.Consul.Host
	p := ViperConf.Consul.Port
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = config.Address
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags

	severAddr := fmt.Sprintf("http://%s:%d/health", host, port)
	check := api.AgentServiceCheck{
		HTTP:     severAddr,
		Timeout:  "3s",
		// 每秒测一次
		Interval: "1s",
		// 5秒不通自动注销
		DeregisterCriticalServiceAfter: "5s",
	}
	agentServiceRegistration.Check = &check

	return client.Agent().ServiceRegister(agentServiceRegistration)
}

// 获取服务列表
func GetServerList() error{
	config := api.DefaultConfig()
	h := ViperConf.Consul.Host
	p := ViperConf.Consul.Port
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	services, err := client.Agent().Services()
	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}

	return nil
}

// 过滤服务
func FilterService() error{
	config := api.DefaultConfig()
	h := ViperConf.Consul.Host
	p := ViperConf.Consul.Port
	config.Address = fmt.Sprintf("%s:%d",h,p)
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	services, err := client.Agent().ServicesWithFilter("Service==accountWeb")
	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--------------------")
	}

	return nil
}