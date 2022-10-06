package biz

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/custom_error"
	"product/internal"
	"product/model"
	"product/proto/pb"
)

func (p ProductServer) CategoryBrandList(ctx context.Context, req *pb.PagingReq) (*pb.CategoryBrandListRes, error) {

	var items []model.ProductCategoryBrand
	var total int64
	var res pb.CategoryBrandListRes
	internal.DB.Model(&model.ProductCategoryBrand{}).Count(&total)
	res.Total = int32(total)
	internal.DB.
		Preload("Category").
		Preload("Brand").
		Scopes(internal.Paginate(int(req.PageNo),int(req.PageSize))).
		Find(&items)

	var cbs []*pb.CategoryBrandRes
	for _, item := range items {
		cbs = append(cbs,ConvertProductCategoryBrand2Pb(item))
	}
	res.Total = int32(total)
	res.ItemList = cbs

	return &res,nil
}

func (p ProductServer) GetCategoryBrandList(ctx context.Context, req *pb.CategoryItemReq) (*pb.BrandRes, error) {
	var res pb.BrandRes
	var category model.Category
	var itemList []model.ProductCategoryBrand
	var itemListRes []*pb.BrandItemRes

	r := internal.DB.First(&category, req.Id)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.CategoryNotExist)
	}
	r = internal.DB.Preload("Brand").Where(&model.ProductCategoryBrand{CategoryId: req.Id}).Find(&itemList)
	if r.RowsAffected > 0 {
		res.Total = int32(r.RowsAffected)
	}
	for _, item := range itemList {
		itemListRes = append(itemListRes, &pb.BrandItemRes{
			Id:   item.Brand.Id,
			Name: item.Brand.Name,
			Logo: item.Brand.Logo,
		})
	}
	res.ItemList = itemListRes
	return &res, nil
}

func (p ProductServer) CreateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	var res pb.CategoryBrandRes
	var item model.ProductCategoryBrand
	var category model.Category
	var brand model.Brand

	// 分类判断
	tx := internal.DB.First(&category, req.CategoryId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}

	// 品牌判断
	tx = internal.DB.First(&brand, req.BrandId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}

	// 是否已经存在关系
	item.CategoryId = req.CategoryId
	item.BrandId = req.BrandId
	internal.DB.Save(&item)

	res.Id = item.Id
	return &res,nil
}

func (p ProductServer) UpdateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {

	var productCategoryBrand model.ProductCategoryBrand
	tx := internal.DB.First(&productCategoryBrand, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.ProductCategoryBrandNotExist)
	}

	if req.CategoryId!=0{
		// 分类判断
		var category model.Category
		tx := internal.DB.First(&category, req.CategoryId)
		if tx.RowsAffected<1{
			return nil,errors.New(custom_error.CategoryNotExist)
		}
		productCategoryBrand.CategoryId = category.Id
	}

	if req.BrandId!=0{
		// 品牌判断
		var brand model.Brand
		tx := internal.DB.First(&brand, req.BrandId)
		if tx.RowsAffected<1{
			return nil,errors.New(custom_error.CategoryNotExist)
		}
		productCategoryBrand.BrandId = brand.Id
	}

	save := internal.DB.Save(&productCategoryBrand)
	fmt.Println(save)

	return &emptypb.Empty{},nil
}

func (p ProductServer) DeleteCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	tx := internal.DB.Delete(&model.ProductCategoryBrand{}, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.DelCategoryBrandFailed)
	}

	return &emptypb.Empty{},nil
}

func ConvertProductCategoryBrand2Pb(pcb model.ProductCategoryBrand) *pb.CategoryBrandRes {
	return &pb.CategoryBrandRes{
		Id: pcb.Id,
		Brand: &pb.BrandItemRes{
			Id:   pcb.Brand.Id,
			Name: pcb.Brand.Name,
			Logo: pcb.Brand.Logo,
		},
		Category: &pb.CategoryItemRes{
			Id:               pcb.Category.Id,
			Name:             pcb.Category.Name,
			ParentCategoryId: pcb.Category.ParentCategoryId,
			Level:            pcb.Category.Level,
		},
	}
}
