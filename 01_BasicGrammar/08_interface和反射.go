package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func demo008() {
	var s Shape = Rectangle{10, 5}

	// 获取变量的类型信息
	t := reflect.TypeOf(s)
	fmt.Println("Type of s:", t)

	// 获取变量的值信息
	v := reflect.ValueOf(s)
	fmt.Println("Value of s:", v)

	// 获取结构体字段名和值
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		//fmt.Printf("%s: %v\n", t.Field(i).Name, v.Interface())
		fmt.Printf("%s: %v\n", t.Field(i).Name, f.Interface())
	}
}
