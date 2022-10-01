package biz

import (
	"account/account_srv/proto/pb"
	"account/internal"
	"context"
	"fmt"
	"testing"
)

func init(){
	internal.InitDB()
}

func TestAccountServer_AddAccount(t *testing.T) {
	server := AccountServer{}
	for i := 0; i < 5; i++ {
		mobile := fmt.Sprintf("1300000000%d",i)
		res, err := server.AddAccount(context.Background(), &pb.AddAccountRequest{
			Mobile:   mobile,
			Password: mobile,
			Nickname: mobile,
			Gender:   "male",
		})
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(res.Id)
	}
}

func TestAccountServer_GetAccountById(t *testing.T) {
	server := AccountServer{}
	for i := 1; i < 5; i++ {
		res, err := server.GetAccountById(context.Background(), &pb.IdRequest{Id: uint32(i)})
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(res)
	}
}

func TestAccountServer_GetAccountList(t *testing.T) {
	server := AccountServer{}
	list, err := server.GetAccountList(context.Background(), &pb.PagingRequest{PageNo: 1, PageSize: 10})
	if err != nil {
		t.Error(err)
	}
	for _, v := range list.AccountList {
		t.Log(v)
	}
}

func TestAccountServer_GetAccountByMobile(t *testing.T) {
	server := AccountServer{}
	for i := 0; i < 6; i++ {
		res, err := server.GetAccountByMobile(context.Background(), &pb.MobileRequest{Mobile: fmt.Sprintf("1300000000%d", i)})
		if err != nil {
			t.Error(err)
		}
		t.Log(res)
	}
}

func TestAccountServer_UpdateAccount(t *testing.T) {
	accountServer := AccountServer{}
	req := pb.UpdateAccountRequest{
		Id:       1,
		Mobile:   "13000000000",
		Password: "123456",
		Nickname: "ffff",
		Gender:   "male",
		Role:     2,
	}
	account, err := accountServer.UpdateAccount(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	t.Log(account.Result)
}


func TestAccountServer_CheckPassword(t *testing.T) {
	accountServer := AccountServer{}
	res, err := accountServer.CheckPassword(context.Background(), &pb.CheckPasswordRequest{
		Password:       "13000000004",
		HashedPassword: "ccedbc711ff2f303d168e5188cc500aa705b97e5597a9e15b13f4799228802c2",
		AccountId:      5,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(res.Result)
}
