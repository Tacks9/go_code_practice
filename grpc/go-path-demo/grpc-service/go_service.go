package main

import (
	"fmt"
	"google.golang.org/grpc"
	"./service"
	"log"
	"net"
)

func main() {

	fmt.Println("【grpc服务端】")

	// 【1】 new 一个 grpc service
	rpcService := grpc.NewServer()

	// 【2】 注册 ProdService
	service.RegisterProdServiceServer(rpcService, new(service.ProdService))

	// 【3】 新创建一个listener 以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 【4】 运行rpcService
	_ = rpcService.Serve(listener)
}
