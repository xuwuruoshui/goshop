package biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"goshop/account_srv/model"
	"goshop/account_srv/proto/pb"
	"goshop/comm"
	"goshop/custom_error"
	"goshop/internal"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

func Paginate(pageNo,pageSize int)func(db *gorm.DB)*gorm.DB{
	return func(db *gorm.DB) *gorm.DB {
		// 默认第一页
		if pageNo==0{
			pageNo=1
		}

		// 最大页数100,默认10
		if pageSize>100{
			pageSize=100
		}else if pageSize<=0{
			pageSize=10
		}

		// 分页
		offset := (pageNo-1)*pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Model2Pb(account model.Account)*pb.AccountRes{
	accountRes := &pb.AccountRes{
		Id:       int32(account.ID),
		Mobile:   account.Mobile,
		Password: account.Password,
		Nickname: account.NickName,
		Gender:   account.Gender,
		Role:     uint32(account.Role),
	}
	return accountRes
}
func (a *AccountServer) GetAccountList(ctx context.Context,req *pb.PagingRequest) (*pb.AccountListRes, error) {

	var accountList []model.Account
	result := internal.DB.Scopes(Paginate(int(req.PageNo),int(req.PageSize))).Find(&accountList)
	if result.Error!=nil{
		return nil,result.Error
	}
	accountListRes := &pb.AccountListRes{}
	accountListRes.Total = int32(result.RowsAffected)

	for _, account := range accountList {
		accountRes := Model2Pb(account)
		accountListRes.AccountList = append(accountListRes.AccountList,accountRes)
	}
	return accountListRes,nil
}
func (a *AccountServer) GetAccountByMobile(ctx context.Context,req *pb.MobileRequest) (*pb.AccountRes, error) {

	var account model.Account
	result := internal.DB.Where(&model.Account{Mobile: req.Mobile}).First(&account)

	if result.RowsAffected==0 {
		return nil,errors.New(custom_error.AccountNotFound)
	}

	res := Model2Pb(account)
	return res,nil
}
func (a *AccountServer) GetAccountById(ctx context.Context,req *pb.IdRequest) (*pb.AccountRes, error) {
	var account model.Account
	result := internal.DB.First(&account,req.Id)

	if result.RowsAffected==0 {
		return nil,errors.New(custom_error.AccountNotFound)
	}

	res := Model2Pb(account)
	return res,nil
}
func (a *AccountServer) AddAccount(ctx context.Context,req *pb.AddAccountRequest) (*pb.AccountRes, error) {

	// 1.判断账户是否存在
	var account model.Account
	tx := internal.DB.Where(&model.Account{Mobile: req.Mobile}).First(&account)
	if tx.RowsAffected==1{
		return nil,errors.New(custom_error.AccountExists)
	}

	// 2.创建
	account.Role=1
	account.Mobile=req.Mobile
	account.NickName=req.Nickname

	salt, hashed := comm.Encode(req.Password)
	account.Password=hashed
	account.Salt=salt
	res := internal.DB.Create(&account)
	if res.Error!=nil{
		return nil,errors.New(custom_error.InternalError)
	}

	accountRes := Model2Pb(account)
	return accountRes,nil
}
func (a *AccountServer) UpdateAccount(ctx context.Context,req *pb.UpdateAccountRequest) (*pb.UpdateAccountRes, error) {
	account := model.Account{}
	tx := internal.DB.First(&account, req.Id)
	if tx.RowsAffected==0{
		return nil,errors.New(custom_error.AccountNotFound)
	}

	// TODO mobile应该是唯一的
	account.Mobile=req.Mobile
	account.NickName=req.Nickname
	account.Gender=req.Gender
	r := internal.DB.Save(&account)
	if r.Error!=nil{
		return nil,errors.New(custom_error.InternalError)
	}
	return &pb.UpdateAccountRes{Result: true},nil
}
func (a *AccountServer) CheckPassword(ctx context.Context,req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	var account model.Account
	tx := internal.DB.First(&account, req.AccountId)

	if tx.Error!=nil{
		return nil,errors.New(custom_error.InternalError)
	}

	if account.Salt==""{
		return nil,errors.New(custom_error.SaltError)
	}

	r := comm.Decode(req.Password, account.Salt, account.Password)

	return &pb.CheckPasswordRes{Result: r},nil
}
