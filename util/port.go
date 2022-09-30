package main

import (
	"fmt"
	"net"
)

func main(){

	// 解析地址
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	// 利用 ListenTCP 方法的如下特性
	// 如果 addr 的端口字段为0，函数将选择一个当前可用的端口
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	// 关闭资源
	defer listener.Close()

	// 为了拿到具体的端口值，我们转换成 *net.TCPAddr类型获取其Port
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Println(port)


}

