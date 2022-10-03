package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/proto/pb"
	"testing"
)

func TestProductServer_CreateCategory(t *testing.T) {
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
		want    *pb.CategoryItemRes
		wantErr bool
	}{
		{args: args{req: &pb.CategoryItemReq{Level: 1, Name: "食物"}}},
		{args: args{req: &pb.CategoryItemReq{Level: 2, Name: "肉类"}}},
		{args: args{req: &pb.CategoryItemReq{Level: 3, Name: "牛肉"}}},
	}

	var id int32 = 1
	for _, tt := range tests {

		p := ProductServer{
			UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
		}
		tt.args.req.ParentCategoryId = id
		got, err := p.CreateCategory(tt.args.ctx, tt.args.req)
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
			//return
		}

		id = got.Id
	}
}

func TestProductServer_DeleteCategory(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoryDelReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.CategoryDelReq{Id: 17}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.DeleteCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductServer_GetAllCategoryList(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx   context.Context
		empty *emptypb.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CategoriesRes
		wantErr bool
	}{
		{args: args{empty: &emptypb.Empty{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.GetAllCategoryList(tt.args.ctx, tt.args.empty)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCategoryList() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(got.CategoryJsonFormat)
		})
	}
}

func TestProductServer_GetSubCategory(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CategoriesReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SubCategoriesRes
		wantErr bool
	}{
		{args: args{req: &pb.CategoriesReq{Id: 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.GetSubCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(got.CategoryJsonFormat)
		})
	}
}

func TestProductServer_UpdateCategory(t *testing.T) {
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
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.CategoryItemReq{Id: 17,Name: "电子数码"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.UpdateCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)

			}
		})
	}
}
