package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"stock/custom_error"
	"stock/internal"
	"stock/model"
	"stock/proto/pb"
)

type StockServer struct {
	pb.UnimplementedStockServiceServer
}

func (s StockServer) SetStock(ctx context.Context, req *pb.ProductStockItem) (*emptypb.Empty, error) {
	var stock model.Stock
	internal.DB.Where("product_id=?",req.ProductId).Find(&stock)
	if stock.Id<1{
		stock.ProductId = req.ProductId
	}
	stock.Num = req.Num

	internal.DB.Save(&stock)
	return &emptypb.Empty{},nil
}

func (s StockServer) StockDetail(ctx context.Context, req *pb.ProductStockItem) (*pb.ProductStockItem, error) {
	var stock model.Stock
	tx := internal.DB.Where("product_id=?", req.ProductId).Find(&stock)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.StockNotFound)
	}

	stockPb := ConvertStockModel2Pb(stock)
	return stockPb,nil
}

func (s StockServer) Sell(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s StockServer) BackStock(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {

	// 还原库存
	// 超时
	// 订单创建失败
	// 手动归还

	tx := internal.DB.Begin()
	for _, item := range req.StockItemList {
		var stock model.Stock
		r := tx.Where("product_id=?", item.ProductId).First(&stock)
		if r.RowsAffected<1{
			tx.Rollback()
			return nil,errors.New(custom_error.StockNotFound)
		}
		stock.Num +=item.Num
		tx.Save(&stock)
	}
	tx.Commit()

	return &emptypb.Empty{},nil
}

func (s StockServer) mustEmbedUnimplementedStockServiceServer() {
	//TODO implement me
	panic("implement me")
}

func ConvertStockModel2Pb(s model.Stock) *pb.ProductStockItem{
	return &pb.ProductStockItem{
		ProductId: s.ProductId,
		Num:       s.Num,
	}
}