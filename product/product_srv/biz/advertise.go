package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/custom_error"
	"product/internal"
	"product/model"
	"product/proto/pb"
)

func (p ProductServer) AdvertiseList(ctx context.Context, empty *emptypb.Empty) (*pb.AdvertiseRes, error) {
	var advertiseList []model.Advertise
	var advertises []*pb.AdvertiseItemRes
	var advertiseRes pb.AdvertiseRes

	r := internal.DB.Find(&advertiseList)
	for _, item := range advertiseList {
		advertises = append(advertises,ConvertAdvertiseModel2Pb(item))
	}

	advertiseRes.ItemList = advertises
	advertiseRes.Total = int32(r.RowsAffected)
	return &advertiseRes,nil
}

func (p ProductServer) CreateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*pb.AdvertiseItemRes, error) {
	var ad model.Advertise
	ad.Index = req.Index
	ad.Url = req.Url
	ad.Image = req.Image
	// 不存在则插入,存在则更新
	internal.DB.Save(&ad)

	return ConvertAdvertiseModel2Pb(ad),nil
}

func (p ProductServer) UpdateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {
	var ad model.Advertise
	tx := internal.DB.First(&ad, req.Id)

	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.AdvertiseNotExist)
	}

	if req.Index>0{
		ad.Index = req.Index
	}

	if req.Image!=""{
		ad.Image = req.Image
	}

	if req.Url!=""{
		ad.Url = req.Url
	}
	internal.DB.Save(&ad)
	return &emptypb.Empty{},nil
}

func (p ProductServer) DeleteAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {

	tx := internal.DB.Delete(&model.Advertise{}, req.Id)
	if tx.Error!=nil{
		return nil,errors.New(custom_error.DeleteAdvertiseFailed)
	}
	return &emptypb.Empty{},nil
}


func ConvertAdvertiseModel2Pb(item model.Advertise)*pb.AdvertiseItemRes{

	return &pb.AdvertiseItemRes{
		Id:    item.Id,
		Index: item.Index,
		Image: item.Image,
		Url:   item.Url,
	}
}
