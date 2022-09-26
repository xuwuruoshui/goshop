package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"goshop/account_srv/proto/pb"
	"goshop/account_web/res"
	"goshop/custom_error"
	"goshop/log"
	"net/http"
	"strconv"
)

func HandleError(err error)string{
	if err!=nil{
		switch err.Error() {
		case custom_error.AccountExists:
			return custom_error.AccountExists
		case custom_error.AccountNotFound:
			return custom_error.AccountNotFound
		case custom_error.SaltError:
			return custom_error.SaltError
		default:
			return custom_error.InternalError
		}
	}
	return ""
}

func AccountListHandler(c *gin.Context){
	pageNo,_ := strconv.ParseInt(c.Query("pageNo"),10,32)
	pageSize,_ := strconv.ParseInt(c.Query("pageSize"),10,32)

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err!=nil{
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)
	list, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	})
	if err!=nil{
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}

	var accountResList []*res.Account4Res
	for _,item := range list.AccountList {
		tmp := item
		accountResList = append(accountResList,pb2Res(tmp))
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
		"total": list.Total,
		"data": accountResList,
	})
}

func pb2Res(accountRes *pb.AccountRes) *res.Account4Res{
	return &res.Account4Res{
		Mobile:   accountRes.Mobile,
		NickName:  accountRes.Nickname,
		Gender:    accountRes.Gender,
	}
}