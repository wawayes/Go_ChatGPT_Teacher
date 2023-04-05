package main

import "fmt"

func demo002() {
	// 创建一个空的字符串切片
	var fruits []string
	// 添加元素
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "watermelon")
	// 打印切片长度和容量
	fmt.Printf("len=%d, cap=%d\n", len(fruits), cap(fruits))
	// 遍历切片并打印元素
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, fruit)
	}
	// 截取切片
	slicedFruits := fruits[1:3]
	fmt.Printf("len=%d, cap=%d\n", len(slicedFruits), cap(slicedFruits))
	// 修改元素
	slicedFruits = append(slicedFruits, "lemon")
	fmt.Printf("len=%d, cap=%d\n", len(slicedFruits), cap(slicedFruits))
	slicedFruits = append(slicedFruits, "pineapple")
	fmt.Printf("len=%d, cap=%d\n", len(slicedFruits), cap(slicedFruits))
	fmt.Println(slicedFruits)
}
