package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"goshop/account_srv/biz"
	"goshop/account_srv/proto/pb"
	"goshop/internal"
	"net"
)

func init(){
	internal.InitDB()
}

func main(){
	ip := flag.String("ip","192.168.0.112","输入ip")
	port := flag.Int("port",9095,"输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d",*ip,*port)

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server,&biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpc_health_v1.RegisterHealthServer(server,health.NewServer())
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d",
		internal.AppConf.Consul.Host,
		internal.AppConf.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	checkAddr := fmt.Sprintf("%s:%d",internal.AppConf.AccountSrv.Host,internal.AppConf.AccountSrv.Port)
	check := &api.AgentServiceCheck{
		GRPC: checkAddr,
		Timeout: "3s",
		Interval: "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	reg := &api.AgentServiceRegistration{
		Name: internal.AppConf.AccountSrv.SrvName,
		ID: internal.AppConf.AccountSrv.SrvName,
		Port: internal.AppConf.AccountSrv.Port,
		Tags: internal.AppConf.AccountSrv.Tags,
		Address: internal.AppConf.AccountSrv.Host,
		Check: check,
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
