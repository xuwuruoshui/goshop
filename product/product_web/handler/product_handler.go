package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"product/custom_error"
	"product/internal"
	"product/proto/pb"
	"strconv"
)

var client pb.ProductServiceClient

func initGRPC() error {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.Consul.Host, internal.AppConf.Consul.Port)
	dialAddr := fmt.Sprintf("consul://%s/productSrv?wait=14", addr)

	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
	}

	client = pb.NewProductServiceClient(conn)
	return nil
}

func init() {
	err := initGRPC()
	if err != nil {
		panic(err)
	}
}

func ProductListHandler(c *gin.Context) {
	var condition pb.ProductConditionReq

	minPriceStr := c.DefaultQuery("minPrice", "0")
	minPrice, err := strconv.Atoi(minPriceStr)
	if err != nil {
		zap.S().Error("minPrice error")
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}

	maxPriceStr := c.DefaultQuery("maxPrice", "0")
	maxPrice, err := strconv.Atoi(maxPriceStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}

	condition.MinPrice = int32(minPrice)
	condition.MaxPrice = int32(maxPrice)

	categoryIdStr := c.DefaultQuery("categoryId", "0")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}

	condition.CategoryId = int32(categoryId)

	isHot := c.DefaultQuery("isPop", "0")
	if isHot == "1" {
		condition.IsPop = true
	}

	isNew := c.DefaultQuery("isNew", "0")
	if isNew == "1" {
		condition.IsNew = true
	}

	pageNoStr := c.DefaultQuery("pageNo", "1")
	pageNo, err := strconv.Atoi(pageNoStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}
	condition.PageNo = int32(pageNo)

	pageSizeStr := c.DefaultQuery("pageSize", "0")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}
	condition.PageSize = int32(pageSize)

	keyword := c.DefaultQuery("keyword", "")
	condition.KeyWord = keyword

	list, err := client.ProductList(context.Background(), &condition)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "产品列表查询失败",
			//默认值
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"total": list.Total,
		"data":  list.ItemList,
	})
}

// 健康检查
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
