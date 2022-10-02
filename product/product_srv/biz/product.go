package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/proto/pb"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) BatchGetProdcut(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateProduct(ctx context.Context, item *pb.CreateProductItem) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteProduct(ctx context.Context, item *pb.ProductDelItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateProduct(ctx context.Context, item *pb.CreateProductItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

