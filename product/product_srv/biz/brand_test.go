package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/internal"
	"product/proto/pb"
	"testing"
)

//var client  pb.ProductServiceClient
//
//func init() {
//	addr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
//	dialAddr := fmt.Sprintf("consul://%s/accountSrv?wait=14",addr)
//
//	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
//	if err != nil {
//		zap.S().Fatal(err)
//	}
//
//	client = pb.NewProductServiceClient(conn)
//}

func init(){
	internal.InitDB()
}


func TestProductServer_BrandList(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.BrandPagingReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.BrandRes
		wantErr bool
	}{
		{args: args{req: &pb.BrandPagingReq{PageNo: 1,PageSize: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.BrandList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrandList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("BrandList() got = %v, want %v", got, tt.want)
			//}
			fmt.Println(got)
		})
	}
}

func TestProductServer_CreateBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.BrandItemReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.BrandItemRes
		wantErr bool
	}{
		{args: args{req:  &pb.BrandItemReq{Name: "OPPO", Logo: "https://oppo.com"}}},
		{args: args{req:  &pb.BrandItemReq{Name: "One+", Logo: "https://one.com"}}},
		{args: args{req:  &pb.BrandItemReq{Name: "魅族", Logo: "https://meizu.com"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.CreateBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBrand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateBrand() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestProductServer_DeleteBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.BrandItemReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.BrandItemReq{Id: 15}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.DeleteBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteBrand() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("DeleteBrand() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestProductServer_UpdateBrand(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.BrandItemReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.BrandItemReq{Id:14,Name: "华为666",Logo: "https://huawei.cn"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.UpdateBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateBrand() error = %v, wantErr %v", err, tt.wantErr)
				// return
			}
		})
	}
}