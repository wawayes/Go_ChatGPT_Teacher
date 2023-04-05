package main

import (
	"fmt"
)

func init() {
	fmt.Println("这里是init函数，在main函数执行前会调用.")
}

func Demo01() {
	fmt.Println("Hello World")
	fmt.Println("这是我的第一个Hello World程序，使用Go语言，并且我想要将ChatGPT作为我在学习Go语言过程中的首席老师")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	var x int = 1
	var y int = 2
	var sum int = x + y
	output := fmt.Sprintf("sum = %d", sum)
	fmt.Println(output)

	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output = fmt.Sprintf("下标索引为[3:6],输出值为: %d", numbers[3:6])
	fmt.Println(output)

	var ages map[string]int = map[string]int{
		"Alice": 12,
		"Bob":   25,
		"Candy": 22,
	}
	output = fmt.Sprintf("Alice的年龄为: %d", ages["Alice"])
	fmt.Println(output)

	ch := make(chan int)
	go sumFunc(numbers[:len(numbers)/2], ch)
	go sumFunc(numbers[len(numbers)/2:], ch)
	a, b := <-ch, <-ch
	fmt.Println(a + b)

	//	调用demo02的MyVar
	fmt.Println(MyVar)
	// 调用demo02的demo02func
	demo02()
}

func sumFunc(numbers []int, ch chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	ch <- sum
}
