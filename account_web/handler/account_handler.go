package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"goshop/account_srv/proto/pb"
	"goshop/account_web/req"
	"goshop/account_web/res"
	"goshop/custom_error"
	"goshop/internal"
	"goshop/jwt_op"
	"goshop/log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	accountSrvHost string
	accountSrvPort int
	client pb.AccountServiceClient
)

func initConsulClient()error{
	// consul grpc
	config := api.DefaultConfig()
	consulAddr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
	config.Address = consulAddr
	consulClient, err := api.NewClient(config)
	if err != nil {
		zap.S().Error("AccountHandler,创建consul client失败:",err.Error())
		return err
	}

	serverList, err := consulClient.Agent().ServicesWithFilter("Service==accountSrv")
	if err != nil {
		zap.S().Error("AccountHandler,consul获取服务列表失败:",err.Error())
		return err
	}
	for _, v := range serverList {
		accountSrvHost = v.Address
		accountSrvPort = v.Port
	}
	return nil
}

func initGRPC() error{
	grpcAddr := fmt.Sprintf("%s:%d",accountSrvHost,accountSrvPort)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)

		return err
	}

	client = pb.NewAccountServiceClient(conn)
	return nil
}

func init(){
	err := initConsulClient()
	if err != nil {
		panic(err)
	}
	err = initGRPC()
	if err != nil {
		panic(err)
	}
}

func HandleError(err error) string {
	if err != nil {
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
func pb2Res(accountRes *pb.AccountRes) *res.Account4Res {
	return &res.Account4Res{
		Mobile:   accountRes.Mobile,
		NickName: accountRes.Nickname,
		Gender:   accountRes.Gender,
	}
}

// 获取账户列表
func AccountListHandler(c *gin.Context) {

	pageNo, _ := strconv.ParseInt(c.Query("pageNo"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 32)

	list, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败:%s", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var accountResList []*res.Account4Res
	for _, item := range list.AccountList {
		tmp := item
		accountResList = append(accountResList, pb2Res(tmp))
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"total": list.Total,
		"data":  accountResList,
	})
}

// 密码登录
func LoginByPasswordHandler(c *gin.Context) {
	var loginByPassword req.LoginByPassword
	err := c.ShouldBindJSON(&loginByPassword)
	if err != nil {
		log.Logger.Error("LoginByPassword出错" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "解析参数错误",
		})
		return
	}

	reg := regexp.MustCompile("^1[345789]{1}\\d{9}$")
	match := reg.Match([]byte(loginByPassword.Mobile))
	if !match {
		log.Logger.Error("LoginByPassword出错" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "手机号格式不正确",
		})
	}

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Logger.Error("LoginByPassword拨号错误" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	client := pb.NewAccountServiceClient(conn)

	account, err := client.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: loginByPassword.Mobile,
	})
	if err != nil {
		log.Logger.Error("GetAccountByMobile 错误" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}

	checkRes, err := client.CheckPassword(context.Background(), &pb.CheckPasswordRequest{
		AccountId:      uint32(account.Id),
		Password:       loginByPassword.Password,
		HashedPassword: account.Password,
	})

	if err != nil {
		log.Logger.Error("CheckPassword 错误" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 登录成功生成token
	checkResult := "登录失败"
	if !checkRes.Result {

		c.JSON(http.StatusOK, gin.H{
			"msg":    "",
			"result": checkResult,
			"token":  "",
		})
		return
	}

	checkResult = "登录成功"
	j := jwt_op.NewJWT()
	now := time.Now()
	claims := jwt_op.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 30)),
		},
		Id:          account.Id,
		NickName:    account.Nickname,
		AuthorityId: int32(account.Role),
	}
	token, err := j.GenerateJWT(claims)
	if err != nil {
		log.Logger.Error("CheckPassword 错误" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "",
		"result": checkResult,
		"token":  token,
	})

}

// 健康检查
func HealthHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"OK",
	})
}