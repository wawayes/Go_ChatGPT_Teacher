package main

import "fmt"

var MyVar string = "可见性：首字母大写则其他包也可调用"

func demo02() {
	fmt.Println("demo02" + MyVar)
}
