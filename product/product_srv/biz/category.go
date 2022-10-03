package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/custom_error"
	"product/internal"
	"product/model"
	"product/proto/pb"
)

func (p ProductServer) GetAllCategoryList(ctx context.Context, empty *emptypb.Empty) (*pb.CategoriesRes, error) {

	var categoryList []model.Category
	internal.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categoryList)
	var res pb.CategoriesRes


	str, err := json.Marshal(&categoryList)
	if err != nil {
		return nil,errors.New(custom_error.CategoryMarshalFailed)
	}

	res.CategoryJsonFormat = string(str)

	return &res,nil
}

func (p ProductServer) GetSubCategory(ctx context.Context, req *pb.CategoriesReq) (*pb.SubCategoriesRes, error) {
	var category model.Category
	var res pb.SubCategoriesRes
	tx := internal.DB.First(&category, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}
	pre := "SubCategory"
	if category.Level==1{
		pre = "SubCategory.SubCategory"
	}
	var subCategoryList []*model.Category
	internal.DB.Where(&model.Category{ParentCategoryId: req.Id}).Preload(pre).Find(&subCategoryList)
	category.SubCategory = subCategoryList

	str, err := json.Marshal(category)
	if err != nil {
		return nil,errors.New(custom_error.CategoryMarshalFailed)
	}
	res.CategoryJsonFormat = string(str)
	return &res,nil
}

func (p ProductServer) CreateCategory(ctx context.Context, req *pb.CategoryItemReq) (*pb.CategoryItemRes, error) {
	category := model.Category{}
	category.Name = req.Name
	category.Level = req.Level
	if category.Level>=1{
		category.ParentCategoryId = req.ParentCategoryId
	}

	tx := internal.DB.Save(&category)
	fmt.Println(tx)
	return ConvertCategoryModel2Pb(category),nil
}

func (p ProductServer) UpdateCategory(ctx context.Context, req *pb.CategoryItemReq) (*emptypb.Empty, error) {
	var category model.Category
	tx := internal.DB.Find(&category, req.Id)
	if tx.RowsAffected<1{
		return nil,errors.New(custom_error.CategoryNotExist)
	}

	if req.Name!=""{
		category.Name = req.Name
	}

	if req.ParentCategoryId>0{
		category.ParentCategoryId = req.ParentCategoryId
	}

	if req.Level>0{
		category.Level = req.Level
	}
	internal.DB.Save(&category)
	return &emptypb.Empty{},nil
}

func (p ProductServer) DeleteCategory(ctx context.Context, req *pb.CategoryDelReq) (*emptypb.Empty, error) {
	internal.DB.Delete(&model.Category{},req.Id)
	// TODO 逻辑判断
	// 如果删除的是一级分类,下面的2级、3级也得删除
	return &emptypb.Empty{},nil
}

func ConvertCategoryModel2Pb(category model.Category)*pb.CategoryItemRes{
	return &pb.CategoryItemRes{
		Id:               category.Id,
		Name:             category.Name,
		ParentCategoryId: category.ParentCategoryId,
		Level:            category.Level,
	}
}