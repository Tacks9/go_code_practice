package main

import "fmt"

func main() {
	fmt.Println("for-range -------------------------for-range 的配套用法")

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[4] = 4.0
	map1[3] = 3.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	/*
		// 两次打印顺序不同 =》 随机
			Map item: Capital of Italy is Rome
			Map item: Capital of Japan is Tokyo
			Map item: Capital of France is Paris

			Map item: Capital of France is Paris
			Map item: Capital of Italy is Rome
			Map item: Capital of Japan is Tokyo
	*/
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

}

/*
for key, value := range map1 {
    // code
}

// 如果你只关心值，可以这么使用
for _, value := range map1 {
    // code
}
// 如果只想获取 key
for key := range map1 {
    fmt.Printf("key is: %d\n", key)
}


1、for-range 的配套用法
2、第一个返回值 key 是 map 中的 key 值，第二个返回值则是该 key 对应的 value 值
2、注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的



*/
