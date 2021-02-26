/*
 * @Descripttion: 【GRPC客户端】
 * @Author: tacks321@qq.com
 * @Date: 2021-02-26 14:42:16
 * @LastEditTime: 2021-02-26 15:00:29
 */
package main

import (
	"context"
	"fmt"
	"grpc-client/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("【grpc 客户端】")

	// [1] new grpc connect
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// [2] defer close
	defer conn.Close()

	// [3] new client
	productServiceClient := service.NewProdServiceClient(conn)

	// [4] call function
	resp, err := productServiceClient.GetProductName(context.Background(), &service.ProductRequest{ProdId: 3})
	if err != nil {
		log.Fatal("调用GRPC方法错误", err)
	}
	// [5] ok
	fmt.Println("调用GRPC方法成功，ProdName=", resp.ProdName)
}
