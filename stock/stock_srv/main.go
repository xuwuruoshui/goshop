package main

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"stock/internal"
	"stock/internal/register"
	"stock/proto/pb"
	"stock/stock_srv/biz"
	"stock/utils"
	"syscall"
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

	randomId := uuid.New().String()
	registry := register.NewConsulRegistry(internal.AppConf.StockSrv.Host, port, register.GRPC)
	err = registry.Register(internal.AppConf.StockSrv.SrvName, randomId, port, internal.AppConf.StockSrv.Tags)
	if err != nil {
		zap.S().Panic("stock GPRC注册失败")
	}

	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()
	q := make(chan os.Signal)
	signal.Notify(q,syscall.SIGINT,syscall.SIGTERM)
	<-q
	err = registry.DeRegister(randomId)
	if err != nil {
		zap.S().Panic("stock注销失败",randomId+":"+err.Error())
	}else{
		zap.S().Info("stock注销成功",randomId)
	}
}
