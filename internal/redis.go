package internal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

var RedisClient *redis.Client

func InitRedis(){
	host := ViperConf.Redis.Host
	port := ViperConf.Redis.Port
	addr := fmt.Sprintf("%s:%d",host,port)
	fmt.Println(addr)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})


	ping := RedisClient.Ping(context.Background())
	fmt.Println(ping.String())
	fmt.Println("Redis 初始化完成")
}