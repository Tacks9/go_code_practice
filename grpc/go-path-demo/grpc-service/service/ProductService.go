package service

import (
	"context"
	"fmt"
)

type ProdService struct {
	
}

// 实现接口的方法
func (ps *ProdService) GetProductName(ctx context.Context,request *ProductRequest) (*ProductResponse,error) {

	// 拿到全部商品
	var Products = [5]string{"苹果","香蕉","橙子","蓝莓","猕猴桃"}
	fmt.Println(Products)

	// 接收参数
	id := int(request.ProdId)
	name:= ""

	// 边界判断
	if id > len(Products) || (id < 0) {
		name = "商品不存在"
	}

	// 根据对应id拿到不同的商品
	name = Products[request.ProdId]

	return &ProductResponse{ProdName: name},nil
}