package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/proto/pb"
	"testing"
)

func TestProductServer_AdvertiseList(t *testing.T) {
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
		want    *pb.AdvertiseRes
		wantErr bool
	}{
		{args: args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := p.AdvertiseList(tt.args.ctx, tt.args.empty)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdvertiseList() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("AdvertiseList() got = %v, want %v", got, tt.want)
			//}
			fmt.Println(got)
		})
	}
}

func TestProductServer_CreateAdvertise(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.AdvertiseReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AdvertiseItemRes
		wantErr bool
	}{
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/a.png",Index: 1,Url: "https://test.com/a.png"}}},
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/b.png",Index: 2,Url: "https://test.com/b.png"}}},
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/c.png",Index: 3,Url: "https://test.com/c.png"}}},
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/d.png",Index: 4,Url: "https://test.com/d.png"}}},
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/e.png",Index: 5,Url: "https://test.com/e.png"}}},
		{args: args{req: &pb.AdvertiseReq{Image: "https://test.com/f.png",Index: 6,Url: "https://test.com/f.png"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.CreateAdvertise(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdvertise() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateAdvertise() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestProductServer_DeleteAdvertise(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.AdvertiseReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.AdvertiseReq{Id: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.DeleteAdvertise(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAdvertise() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("DeleteAdvertise() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestProductServer_UpdateAdvertise(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.AdvertiseReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{args: args{req: &pb.AdvertiseReq{Id: 2,Url: "https://ffff.cn/test"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductServer{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			_, err := p.UpdateAdvertise(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdvertise() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("UpdateAdvertise() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
