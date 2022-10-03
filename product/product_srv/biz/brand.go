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

func (p ProductServer) BrandList(ctx context.Context, req *pb.BrandPagingReq) (*pb.BrandRes, error) {
	var brandList []model.Brand
	var brands []*pb.BrandItemRes
	var brandRes pb.BrandRes

	// 一次查完
	//r := internal.DB.Find(&brandList)
	//for _, item := range brandList {
	//	brands = append(brands,ConvertBrandModel2Pb(item))
	//}
	//
	//brandRes.ItemList = brands
	//brandRes.Total = int32(r.RowsAffected)
	//return &brandRes,nil

	// 分页查询
	internal.DB.Scopes(internal.Paginate(int(req.PageNo), int(req.PageSize))).Find(&brandList)
	for _, item := range brandList {
		brands = append(brands,ConvertBrandModel2Pb(item))
	}

	var count int64
	internal.DB.Model(&model.Brand{}).Count(&count)

	brandRes.ItemList = brands
	brandRes.Total = int32(count)
	return &brandRes,nil
}

func (p ProductServer) CreateBrand(ctx context.Context, req *pb.BrandItemReq) (*pb.BrandItemRes, error) {
	var brand model.Brand
	tx := internal.DB.Where("name=? and logo=?", req.Name, req.Logo).First(&brand)
	if tx.RowsAffected>0{
		return nil,errors.New(custom_error.BrandExist)
	}
	brand.Name = req.Name
	brand.Logo = req.Logo
	internal.DB.Create(&brand)

	return ConvertBrandModel2Pb(brand),nil
}

func (p ProductServer) UpdateBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	var brand model.Brand
	tx := internal.DB.First(&brand, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.BrandNotExist)
	}
	if req.Name!=""{
		brand.Name = req.Name
	}
	if req.Logo!=""{
		brand.Logo = req.Logo
	}
	internal.DB.Save(&brand)

	return &emptypb.Empty{},nil
}

func (p ProductServer) DeleteBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	tx := internal.DB.Delete(&model.Brand{}, req.Id)
	if tx.Error!=nil{
		return nil,errors.New(custom_error.DelBrandFailed)
	}
	return &emptypb.Empty{},nil
}


func ConvertBrandModel2Pb(item model.Brand)*pb.BrandItemRes{

	return &pb.BrandItemRes{
		Id: item.Id,
		Name: item.Name,
		Logo: item.Logo,
	}
}
