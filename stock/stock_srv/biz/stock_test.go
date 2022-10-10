package biz

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"stock/internal"
	"stock/proto/pb"
	"testing"
)

var client pb.StockServiceClient

func initGRPC() error {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
	dialAddr := fmt.Sprintf("consul://%s/stockSrv?wait=14", addr)

	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
	}

	client = pb.NewStockServiceClient(conn)
	return nil
}

func init() {
	err := initGRPC()
	if err != nil {
		panic(err)
	}
}

func TestStockServer_SetStock(t *testing.T) {
	_, err := client.SetStock(context.Background(), &pb.ProductStockItem{
		ProductId: 1,
		Num:       2,
	})

	if err != nil {
		t.Fatal(err)
	}



}

func TestStockServer_BackStock(t *testing.T) {

}

func TestStockServer_Sell(t *testing.T) {

}

func TestStockServer_StockDetail(t *testing.T) {

}
