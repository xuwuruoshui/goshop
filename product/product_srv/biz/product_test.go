package biz

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"product/internal"
	"product/proto/pb"
	"testing"
)

var client pb.ProductServiceClient


func initGRPC() error{
	addr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
	dialAddr := fmt.Sprintf("consul://%s/productSrv?wait=14",addr)

	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
	}

	client = pb.NewProductServiceClient(conn)
	return nil
}

func init(){
	err := initGRPC()
	if err != nil {
		panic(err)
	}
}


func TestProductServer_BatchGetProduct(t *testing.T) {
	ids := []int32{2, 3, 4}
	res, err := client.BatchGetProduct(context.Background(),
		&pb.BatchProductIdReq{
			Ids: ids,
		})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_CreateProduct(t *testing.T) {
	for i := 0; i < 8; i++ {
		res, err := client.CreateProduct(context.Background(), &pb.CreateProductItem{
			Name:        fmt.Sprintf("黄金牛排%d", i),
			Sn:          "123456789",
			CategoryId:  22,
			Price:       359.00,
			RealPrice:   199.00,
			ShortDesc:   "",
			ProductDesc: "",
			Images:      nil,
			DescImages:  nil,
			CoverImage:  "https://space.bilibili.com/375038855",
			IsNew:       true,
			IsPop:       true,
			Selling:     true,
			BrandId:     18,
			FavNum:      6666,
			SoldNum:     5432,
			IsShipFree:  false,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	}
}

func TestProductServer_DeleteProduct(t *testing.T) {
	client.DeleteProduct(context.Background(),
		&pb.ProductDelItem{
			Id: 8,
		})
}

func TestProductServer_GetProductDetail(t *testing.T) {

}

func TestProductServer_ProductList(t *testing.T) {
	res, err := client.ProductList(context.Background(), &pb.ProductConditionReq{
		PageNo:   2,
		PageSize: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.ItemList)
}

func TestProductServer_UpdateProduct(t *testing.T) {
	client.UpdateProduct(context.Background(), &pb.CreateProductItem{
		Id:         8,
		CategoryId: 17,
		BrandId:    19,
		Name:       "战斧牛排66666",
	})
}
