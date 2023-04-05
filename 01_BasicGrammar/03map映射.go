package main

import (
	"fmt"
	"github.com/k0kubun/pp"
)

func demo003() {
	m := make(map[string]int)
	fmt.Printf("m:%s\n", m)

	//增
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Printf("m:%s\n", m)
	fmt.Println(m)

	//删
	delete(m, "two")
	//改
	m["three"] = 6
	//查
	for key, val := range m {
		fmt.Printf("key=%s, val=%d\n", key, val)
	}
	_, err := pp.Println(m)
	if err != nil {
		return
	}
}
