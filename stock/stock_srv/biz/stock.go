package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"stock/proto/pb"
)

type StockServer struct {
	pb.UnimplementedStockServiceServer
}

func (s StockServer) SetStock(ctx context.Context, item *pb.ProductStockItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s StockServer) StockDetail(ctx context.Context, item *pb.ProductStockItem) (*pb.ProductStockItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s StockServer) Sell(ctx context.Context, item *pb.SellItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s StockServer) BackStock(ctx context.Context, item *pb.SellItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s StockServer) mustEmbedUnimplementedStockServiceServer() {
	//TODO implement me
	panic("implement me")
}

