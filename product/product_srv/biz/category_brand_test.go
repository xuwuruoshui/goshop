package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/proto/pb"
	"testing"
)

func TestProductServer_CategoryBrandList(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.PagingReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CategoryBrandListRes
		wantErr bool
	}{
		{args: args{req: &pb.PagingReq{PageNo: 1,PageSize: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.CategoryBrandList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryBrandList() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Log(got)
		})
	}
}

func TestProductServer_CreateCategoryBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoryBrandReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CategoryBrandRes
		wantErr bool
	}{
		{args: args{req: &pb.CategoryBrandReq{CategoryId: 18,BrandId: 16}}},
		{args: args{req: &pb.CategoryBrandReq{CategoryId: 19,BrandId: 17}}},
		{args: args{req: &pb.CategoryBrandReq{CategoryId: 19,BrandId: 19}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.CreateCategoryBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCategoryBrand() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(got.Id)
		})
	}
}

func TestProductServer_DeleteCategoryBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoryBrandReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.CategoryBrandReq{Id: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.DeleteCategoryBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCategoryBrand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductServer_GetCategoryBrandList(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoryItemReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.BrandRes
		wantErr bool
	}{
		{args: args{req: &pb.CategoryItemReq{Id:18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.GetCategoryBrandList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCategoryBrandList() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(got)
		})
	}
}

func TestProductServer_UpdateCategoryBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoryBrandReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.CategoryBrandReq{Id:2,CategoryId: 19,BrandId: 19}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.UpdateCategoryBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCategoryBrand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
