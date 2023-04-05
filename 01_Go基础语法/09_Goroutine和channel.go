package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(c chan<- int) {
	for {
		n := rand.Intn(100)                                           // 生成0-99之间的随机数
		c <- n                                                        // 发送到通道中
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 随机等待一段时间
	}
}

func consumer(c <-chan int, done chan<- bool) {
	sum := 0
	count := 0
	for n := range c { // 从通道中循环接收随机数
		sum += n
		count++
		fmt.Printf("Received %d, Current Average: %.2f\n", n, float64(sum)/float64(count))
	}
	done <- true
}

func demo009() {
	c := make(chan int)
	done := make(chan bool)

	// 启动生产者和消费者goroutine
	go producer(c)
	go consumer(c, done)

	// 等待消费者goroutine执行完毕
	<-done
}
