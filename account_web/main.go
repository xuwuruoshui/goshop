package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"goshop/account_web/handler"
	"goshop/internal"
)

func init(){
	randUUID := uuid.New().String()
	err := internal.Reg(internal.AppConf.AccountWeb.Host,
		internal.AppConf.AccountWeb.SrvName,
		randUUID,
		internal.AppConf.AccountWeb.Port,
		internal.AppConf.AccountWeb.Tags)
	if err != nil {
		panic(err)
	}
}

func main()  {
	ip := flag.String("ip","192.168.0.112","输入IP")
	port := flag.Int("port",8081,"输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	r := gin.Default()
	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list",handler.AccountListHandler)
		accountGroup.POST("/login",handler.LoginByPasswordHandler)
		accountGroup.GET("/captcha",handler.CaptchaHandler)
	}
	r.GET("/health",handler.HealthHandler)
	r.Run(addr)
}


