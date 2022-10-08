package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"stock/internal"
	"stock/proto/pb"
	"stock/stock_srv/biz"
	"stock/utils"
)

func init(){
	internal.InitDB()
}

func main(){
	port := utils.GenRandPort()
	addr := fmt.Sprintf("%s:%d",internal.AppConf.StockSrv.Host,port)

	server := grpc.NewServer()
	pb.RegisterStockServiceServer(server,&biz.StockServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// 1、GRPC注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server,health.NewServer())

	// 2.创建Consul客户端
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d",
		internal.AppConf.Consul.Host,
		internal.AppConf.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 3.Consul设置GRPC健康检查地址
	// 超时3秒,1秒检测一次,超过5秒注销
	checkAddr := fmt.Sprintf("%s:%d",internal.AppConf.StockWeb.Host,port)
	check := &api.AgentServiceCheck{
		GRPC: checkAddr,
		Timeout: "3s",
		Interval: "1s",
		DeregisterCriticalServiceAfter: "5s",
	}

	// 4.在Consul上注册服务
	randUUID := uuid.New().String()
	reg := &api.AgentServiceRegistration{
		// 服务名称(可以相同，但id必须不一样)
		Name: internal.AppConf.StockSrv.SrvName,
		// 每个实例的id
		ID: randUUID,
		Port: port,
		Tags: internal.AppConf.StockSrv.Tags,
		Address: internal.AppConf.StockSrv.Host,
		Check: check,
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}

	fmt.Println("当前节点端口:",port)
	fmt.Println("当前节点ID:",randUUID)

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
