package main

/**
prompt
给我讲讲go的基础结构和类型，包括数据类型，常量，变量，bool,数值类型,string,错误类型等等，
任你发挥，只要你觉得我必须学习的内容能让我更好的掌握go语言，你就可以直接讲
*/

import (
	"errors"
	"fmt"
)

const (
	pi      = 3.1415926
	version = "v1.0.0"
)

var (
	name string = "GoLang"
	age  int    = 25
)

func demo001() {
	fmt.Println("Hello, " + name)
	fmt.Println("Age:", age)

	// Boolean type
	isTrue := true
	fmt.Println("isTrue:", isTrue)

	// Numeric types
	var a int = 20
	var b int64 = 100
	var c float32 = 3.14
	var d complex64 = 1 + 2i

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// String type
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + " " + str2)

	// Error type
	err := errors.New("Some error occurred")
	if err != nil {
		fmt.Println(err.Error())
	}
}
