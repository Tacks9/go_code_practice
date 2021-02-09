/*
 * @Descripttion: slices 切片，比较常用，引用类型
 * @Author: tacks321@qq.com
 * @Date: 2021-02-08 18:34:26
 * @LastEditTime: 2021-02-09 11:43:26
 */

package main

import "fmt"

func main() {

	fmt.Println("==========================slices切片=================================")

	// =======================【1】 切片依赖于基础数组的创建
	// 创建数组
	var baseArr [6]int
	// 创建切片，容量是3 同时引用数组 3 4 的值，实际上切片是共享了数组的数据
	// 切片在内存中的组织方式 是3 个域的结构体：指向相关数组的指针pointer、切片长度len、切片容量cap
	// 切片本身就是引用类型
	var baseSlice []int = baseArr[3:5]

	// 数组初始化赋值
	for i := 0; i < len(baseArr); i++ {
		baseArr[i] = i
	}

	// 数组打印
	fmt.Println("[baseArr] Init: ", baseArr, " Len: ", len(baseArr))
	// 切片打印  不等式永远成立：0 <= len(slice) <= cap(slice)
	fmt.Println("[baseSlice] Init: ", baseSlice, " Len: ", len(baseSlice), "Cap: ", cap(baseSlice))
	// 切片追加元素
	baseSlice = append(baseSlice, 0)
	fmt.Println("[baseSlice] Init: ", baseSlice, " Len: ", len(baseSlice), "Cap: ", cap(baseSlice))

	// =======================【2】 切片采用make创建，同时创建好基础相关数组，切片大小同数组大小一样

	// 创建一个长度为3，容量也为3的 string类型的 切片slice，并初始化为空字符串
	slice1 := make([]string, 3)
	fmt.Println("[slice1] Init: ", slice1)
	// 设置值
	slice1[0] = "a"
	slice1[1] = "b"
	slice1[2] = "c"
	fmt.Println("[slice1] Set: ", slice1)
	// 获取值
	fmt.Println("[slice1] Get:", slice1[0], slice1[1], slice1[2])

	// 获取切片长度 len()
	fmt.Println(len(slice1))

	// 追加新的元素 append()
	slice1 = append(slice1, "d")
	slice1 = append(slice1, "e", "f")
	fmt.Println("[slice1] Append:", slice1)

	// 获得一个新的切片 左闭右开
	slice2 := slice1[2:4]
	fmt.Println("[slice2] ", slice2)

	// 相当于直接复制一个新的切片
	slice3 := slice1[:]
	fmt.Println("[slice3] ", slice3)

	// 相当于[0,4)范围的slice1的元素 重新组成一个切片
	slice4 := slice1[:4]
	fmt.Println("[slice4] ", slice4)

	// 相当于[5,5)范围的slice1的元素，重新组成一个切片
	slice5 := slice1[5:]
	fmt.Println("[slice5] ", slice5)

	// 直接声明并初始化
	slice6 := []string{"T", "a", "c", "k", "s", "TACKS"}
	fmt.Println("[slice6] Init and Assign: ", slice6)

	// 几种创建切片的方式
	// make创建的是引用类型
	sliceMake1 := make([]int, 10)
	sliceMake2 := make([]int, 5, 10)
	// 先new出来数组，然后获取切片
	sliceNew1 := new([10]int)[0:5]
	fmt.Println("[sliceMake1] ", len(sliceMake1), cap(sliceMake1))
	fmt.Println("[sliceMake2] ", len(sliceMake2), cap(sliceMake2))
	fmt.Println("[sliceNew1] ", len(sliceNew1), cap(sliceNew1))

	// 那么new 和 make 在创建切片的区别是
	// make(T) 适用于引用类型，如 slice\map\channel
	// new(T) 是为T分配内存，返回类型 *T的内存地址 如 array\结构体

	// 二维切片
	sliceTwo := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// 二维切片内部 声明每一个切片元素内存
		sliceTwo[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			sliceTwo[i][j] = i + j
		}
	}
	fmt.Println("[sliceTwo] Init and assign:", sliceTwo)

	// ============================================== 【3】 切片的浅拷贝 、深拷贝

	// 创建一个切片
	sliceCopy1 := make([]string, 5, 5)
	fmt.Println("[sliceCopy1] Init: ", sliceCopy1)
	// 浅拷贝
	sliceCopy2 := sliceCopy1
	fmt.Println("[sliceCopy2] Copy: ", sliceCopy1)

	// 更改其中一个切片，另一个也会改变，相当于其实一样，就是名字不一样
	sliceCopy1[1] = "Tacks"
	fmt.Println("[sliceCopy1]: ", sliceCopy1)
	fmt.Println("[sliceCopy2]: ", sliceCopy2)

	// 创建一个切片
	sliceCopy3 := make([]string, 5, 5)
	sliceCopy3[0] = "Tacks"
	sliceCopy4 := make([]string, 4, 4)
	sliceCopy5 := make([]string, 5, 5)
	sliceCopy6 := make([]string, 6, 6)
	// 拷贝 copy(目标，源值)
	fmt.Println(copy(sliceCopy4, sliceCopy3)) // 4 显然要拷贝的切片要大于目标切片，那么就后面切断
	fmt.Println(copy(sliceCopy5, sliceCopy3)) // 5
	fmt.Println(copy(sliceCopy6, sliceCopy3)) // 5
	// 单独修改
	sliceCopy4[1] = "444444444"
	sliceCopy5[1] = "555555555"
	sliceCopy6[1] = "666666666"
	fmt.Println("[sliceCopy4]: ", sliceCopy4, len(sliceCopy4), cap(sliceCopy4))
	fmt.Println("[sliceCopy5]: ", sliceCopy5, len(sliceCopy5), cap(sliceCopy5))
	fmt.Println("[sliceCopy6]: ", sliceCopy6, len(sliceCopy6), cap(sliceCopy6))

}
