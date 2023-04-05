package main

import "fmt"

// Person 定义一个结构体
type Person struct {
	name string
	age  int
}

// SayHello 定义结构体方法，使用值接收者
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old.\n", p.name, p.age)
}

// SetAge 定义结构体方法，使用指针接收者
func (p *Person) SetAge(age int) {
	p.age = age
}

func demo006() {
	// 创建一个结构体对象
	person := Person{name: "Tom", age: 20}

	// 调用结构体方法
	person.SayHello() // 输出：Hello, my name is Tom, I am 20 years old.

	// 调用指针接收者的方法
	person.SetAge(30)
	person.SayHello() // 输出：Hello, my name is Tom, I am 30 years old.
}
