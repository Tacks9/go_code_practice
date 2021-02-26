package main

import (
	"./service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("【grpc客户端】")

	// 【1】 创建新的连接
	conn,err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// 【2】 主程序退出时 关闭连接
	defer conn.Close()

	// 【3】 调用Product.pb.go中的客户端方法
	productServiceClient := service.NewProdServiceClient(conn)

	// 【4】 直接调用服务端的方法，像在本地一样
	resp, err := productServiceClient.GetProductName(context.Background(), &service.ProductRequest{ProdId: 3})
	if err != nil {
		log.Fatal("调用GRPC方法错误",err)
	}

	// 【5】 调用grpc方法成功
	fmt.Println("调用GRPC方法成功，ProdName=",resp.ProdName)


}
