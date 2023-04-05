package main

import "fmt"

func demo005() {
	// 匿名函数
	add := func(x, y int) int {
		return x + y
	}
	// 将函数作为变量
	fmt.Println(add(3, 4))
	// swap函数
	fmt.Println(swap("hello", "world"))
}

// 两个返回值
func swap(x, y string) (string, string) {
	return y, x
}
