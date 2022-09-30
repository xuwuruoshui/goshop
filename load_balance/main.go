package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"goshop/account_srv/proto/pb"
	"goshop/internal"
)

func main(){

	addr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
	dialAddr := fmt.Sprintf("consul://%s/accountSrv?wait=14",addr)

	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAccountServiceClient(conn)

	for i := 0; i < 10; i++ {
		list, err := client.GetAccountList(context.Background(), &pb.PagingRequest{

			PageNo:   1,
			PageSize: 3,
		})
		if err != nil {
			zap.S().Fatal(err)
		}

		for index, item := range list.AccountList {
			fmt.Println(index,"------------",item)
		}
	}

}
