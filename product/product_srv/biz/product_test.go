package biz

import (
	"context"
	"product/model"
	"product/proto/pb"
	"reflect"
	"testing"
)

func TestConvertProductModel2Pb(t *testing.T) {
	type args struct {
		pro model.Product
	}
	tests := []struct {
		name string
		args args
		want *pb.ProductItemRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertProductModel2Pb(tt.args.pro); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertProductModel2Pb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertReq2Model(t *testing.T) {
	type args struct {
		req      *pb.CreateProductItem
		category *model.Category
		brand    *model.Brand
	}
	tests := []struct {
		name string
		args args
		want *model.Product
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertReq2Model(tt.args.req, tt.args.category, tt.args.brand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertReq2Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_BatchGetProduct(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.BatchProductIdReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ProductRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.BatchGetProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchGetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchGetProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_CreateProduct(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateProductItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ProductItemRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.CreateProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_DeleteProduct(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ProductDelItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.DeleteProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_GetProductDetail(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ProductItemReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ProductItemRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.GetProductDetail(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_ProductList(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.ProductConditionReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ProductRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.ProductList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServer_UpdateProduct(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateProductItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.UpdateProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
