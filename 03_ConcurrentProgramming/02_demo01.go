package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		// 2 秒后往通道 ch 中写入一个数据
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	fmt.Println("waiting...")
	// 等待从通道 ch 中读取数据
	<-ch
	fmt.Println("done!")
}
