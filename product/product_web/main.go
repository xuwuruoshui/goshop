package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"product/internal"
	"product/internal/register"
	"product/product_web/handler"
	"product/utils"
	"syscall"
)

var (
	consulRegistry register.ConsulRegistry
	randomId string
)

func init(){
	randomPort := utils.GenRandPort()
	if !internal.AppConf.Debug{
		internal.AppConf.ProductWeb.Port = randomPort
	}

	randomId = uuid.New().String()
	consulRegistry := register.NewConsulRegistry(internal.AppConf.ProductWeb.Host, internal.AppConf.ProductWeb.Port)
	err := consulRegistry.Register(internal.AppConf.ProductWeb.SrvName, randomId, internal.AppConf.ProductWeb.Port, internal.AppConf.ProductWeb.Tags)
	if err != nil {
		zap.S().Panic(err)
	}

}

func main()  {
	ip := internal.AppConf.ProductWeb.Host
	port := utils.GenRandPort()

	if internal.AppConf.Debug{
		port = internal.AppConf.ProductWeb.Port
	}

	addr := fmt.Sprintf("%s:%d", ip, port)
	r := gin.Default()
	group := r.Group("/v1/product")
	{
		group.GET("/list",handler.ProductListHandler)
		group.GET("/:id",handler.DetailHandler)
		group.POST("/",handler.AddHandler)
		group.PUT("/",handler.UpdateHandler)
		group.DELETE("/:id",handler.DelHandler)
	}
	r.GET("/health",handler.HealthHandler)
	go func() {
		err := r.Run(addr)
		if err != nil {
			zap.S().Panic(addr+"启动失败"+err.Error())
		}else{
			zap.S().Info(addr+"启动成功")
		}
	}()

	q := make(chan os.Signal)
	signal.Notify(q,syscall.SIGINT,syscall.SIGTERM)
	<-q
	err := consulRegistry.DeRegister(randomId)
	if err != nil {
		zap.S().Panic("注销失败",randomId+":"+err.Error())
	}else{
		zap.S().Info("注销成功",randomId)
	}
}


