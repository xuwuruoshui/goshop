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

type ProductServer struct {
	pb.UnimplementedProductServiceServer
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductRes, error) {

	condition := internal.DB.Model(model.Product{})
	var productList []model.Product
	var itemList []*pb.ProductItemRes
	var res pb.ProductRes

	if req.IsPop{
		condition = condition.Where("is_pop=?",req.IsPop)
	}

	if req.IsNew{
		condition = condition.Where("is_new=?",req.IsNew)
	}

	if req.BrandId>0{
		condition = condition.Where("brand_id",req.BrandId)
	}

	if req.KeyWord!=""{
		condition = condition.Where("key_word like ?","%"+req.KeyWord+"%")
	}

	if req.MinPrice>0{
		condition = condition.Where("min_price > ?",req.MinPrice)
	}

	if req.MaxPrice>0{
		condition = condition.Where("max_price>?",req.MaxPrice)
	}

	// 获取三级分类
	if req.CategoryId>0{
		var category model.Category
		tx := internal.DB.First(&category, req.CategoryId)
		if tx.RowsAffected==0{
			return nil,errors.New(custom_error.CategoryNotExist)
		}
		var q string
		if category.Level==1{
			q = "select id from category where parent_category_id in (select id form category where parent_category_id	 = ?)"
		}else if category.Level==2{
			q = "select id from category where parent_category_id = ?"
		}else if category.Level==3{
			q = "select id from category where id = ?"
		}

		condition = condition.Where(fmt.Sprintf("category_id in %s",q),req.CategoryId)
	}

	var count int64
	condition.Count(&count)
	fmt.Println(count)

	condition.Joins("Category").Joins("Brand").Scopes(internal.Paginate(int(req.PageNo),int(req.PageSize))).Find(&productList)
	for _, item := range productList {
		product := ConvertProductModel2Pb(item)
		itemList = append(itemList,product)
	}
	res.Total = int32(count)
	res.ItemList = itemList

	return &res,nil
}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductRes, error) {
	var productList []model.Product
	var res pb.ProductRes
	tx := internal.DB.Find(&productList, req.Ids)
	res.Total = int32(tx.RowsAffected)
	for _, item := range productList {
		pro := ConvertProductModel2Pb(item)
		res.ItemList = append(res.ItemList,pro)
	}

	return &res,nil
}

func (p ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductItem) (*pb.ProductItemRes, error) {

	var category model.Category
	var brand model.Brand
	var res *pb.ProductItemRes

	tx := internal.DB.First(&category, req.CategoryId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}

	tx = internal.DB.First(&brand,req.BrandId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.BrandNotExist)
	}

	var pro model.Product
	ConvertReq2Model(&pro,req, &category, &brand)
	internal.DB.Save(&pro)

	res = ConvertProductModel2Pb(pro)

	return res,nil
}

func (p ProductServer) DeleteProduct(ctx context.Context, req *pb.ProductDelItem) (*emptypb.Empty, error) {
	tx := internal.DB.Delete(&model.Product{}, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.DelProductFailed)
	}
	return &emptypb.Empty{},nil
}

func (p ProductServer) UpdateProduct(ctx context.Context, req *pb.CreateProductItem) (*emptypb.Empty, error) {

	var pro model.Product
	var c model.Category
	var b model.Brand

	tx := internal.DB.First(&pro,req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.ProductNoExist)
	}
	tx = internal.DB.First(&c, req.CategoryId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}
	tx = internal.DB.First(&b, req.BrandId)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.BrandNotExist)
	}

	ConvertReq2Model(&pro,req, &c, &b)
	tx = internal.DB.Save(&pro)
	fmt.Println(tx)
	return &emptypb.Empty{},nil
}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	var pro model.Product
	var res *pb.ProductItemRes
	tx := internal.DB.First(&pro, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.ProductNoExist)
	}

	res = ConvertProductModel2Pb(pro)
	return res,nil
}

func  ConvertReq2Model(p *model.Product, req *pb.CreateProductItem, category *model.Category, brand *model.Brand) *model.Product {

	if req.Id > 0 {
		p.Id = req.Id
	}
	if req.CategoryId > 0 {
		p.CategoryId = req.CategoryId
		p.Category = category
	}
	if req.BrandId > 0 {
		p.BrandId = req.BrandId
		p.Brand = brand
	}
	if req.Selling {
		p.Selling = true
	} else {
		p.Selling = false
	}
	if req.Selling {
		p.Selling = true
	} else {
		p.Selling = false
	}
	if req.IsShipFree {
		p.IsShipFree = true
	} else {
		p.IsShipFree = false
	}
	if req.IsPop {
		p.IsPop = true
	} else {
		p.IsPop = false
	}

	if req.IsNew {
		p.IsNew = true
	} else {
		p.IsNew = false
	}
	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Sn != "" {
		p.SN = req.Sn
	}
	if req.FavNum > 0 {
		p.FavNum = req.FavNum
	}
	if req.SoldNum > 0 {
		p.SoldNum = req.SoldNum
	}
	if req.Price > 0 {
		p.Price = req.Price
	}
	if req.RealPrice > 0 {
		p.RealPrice = req.RealPrice
	}
	if req.ShortDesc != "" {
		p.ShortDesc = req.ShortDesc
	}
	if req.Images != nil {
		p.Images = req.Images
	}
	if req.DescImages != nil {
		p.DescImages = req.DescImages
	}
	if req.CoverImage != "" {
		p.CoverImage = req.CoverImage
	}

	return p
}

func ConvertProductModel2Pb(pro model.Product) *pb.ProductItemRes {

	pbp := &pb.ProductItemRes{
		Id:         pro.Id,
		CategoryId: pro.CategoryId,
		Name:       pro.Name,
		Sn:         pro.SN,
		SoldNum:    pro.SoldNum,
		FavNum:     pro.FavNum,
		Price:      pro.Price,
		RealPrice:  pro.RealPrice,
		ShortDesc:  pro.ShortDesc,
		Images:     pro.Images,
		DescImages: pro.DescImages,
		CoverImage: pro.CoverImage,
		IsNew:      pro.IsNew,
		IsPop:      pro.IsPop,
		Selling:    pro.Selling,
	}

	if pbp.Category!=nil{
		pbp.Category = &pb.CategoryShortItemRes{
			Id:   pro.Category.Id,
			Name: pro.Category.Name,
		}
	}

	if pbp.Brand!=nil{
		pbp.Brand = &pb.BrandItemRes{
			Id:   pro.Brand.Id,
			Name: pro.Brand.Name,
			Logo: pro.Brand.Logo,
		}
	}

	return pbp
}







