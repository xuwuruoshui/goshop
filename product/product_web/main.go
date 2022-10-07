package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"product/internal"
	"product/product_web/handler"
	"product/utils"
)

func init(){
	randUUID := uuid.New().String()
	err := internal.Reg(internal.AppConf.ProductWeb.Host,
		internal.AppConf.ProductWeb.SrvName,
		randUUID,
		internal.AppConf.ProductWeb.Port,
		internal.AppConf.ProductWeb.Tags)
	if err != nil {
		panic(err)
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
	r.Run(addr)
}


