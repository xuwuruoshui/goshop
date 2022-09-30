package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"goshop/account_srv/biz"
	"goshop/account_srv/proto/pb"
	"goshop/internal"
	"goshop/utils"
	"net"
)

func init(){
	internal.InitDB()
}

func main(){
	//ip := flag.String("ip","192.168.0.112","输入ip")
	//port := flag.Int("port",9095,"输入端口")
	//flag.Parse()
	//addr := fmt.Sprintf("%s:%d",*ip,*port)

	port := utils.GenRandPort()
	addr := fmt.Sprintf("%s:%d",internal.AppConf.AccountSrv.Host,port)

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
	checkAddr := fmt.Sprintf("%s:%d",internal.AppConf.AccountSrv.Host,port)
	check := &api.AgentServiceCheck{
		GRPC: checkAddr,
		Timeout: "3s",
		Interval: "1s",
		DeregisterCriticalServiceAfter: "5s",
	}

	randUUID := uuid.New().String()
	reg := &api.AgentServiceRegistration{
		// 服务名称(可以相同，但id必须不一样)
		Name: internal.AppConf.AccountSrv.SrvName,
		// 每个实例的id
		ID: randUUID,
		Port: port,
		Tags: internal.AppConf.AccountSrv.Tags,
		Address: internal.AppConf.AccountSrv.Host,
		Check: check,
	}

	fmt.Println("当前节点端口:",port)
	fmt.Println("当前节点ID:",randUUID)

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
