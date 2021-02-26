/*
 * @Descripttion: 【GRPC服务端】
 * @Author: tacks321@qq.com
 * @Date: 2021-02-26 14:42:16
 * @LastEditTime: 2021-02-26 14:55:23
 */
package main

import (
	"fmt"
	"log"
	"net"

	"grpc-service/service"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("【Grpc 服务端】")

	// 【1】new grpc service
	rpcService := grpc.NewServer()

	// 【2】register Prodservice
	service.RegisterProdServiceServer(rpcService, new(service.ProdService))

	// 【3】new listener
	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatal("服务监听失败", err)
	}

	// 【4】run rpcService
	_ = rpcService.Serve(listener)
}
