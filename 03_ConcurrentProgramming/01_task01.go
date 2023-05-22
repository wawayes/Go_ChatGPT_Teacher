package main

import (
	"fmt"
	"time"
)

func task01() {
	NumJobs := 10
	jobs := make(chan int, NumJobs)
	res := make(chan int, NumJobs)
	for i := 1; i <= 3; i++ {
		go Worker(i, jobs, res)
	}

	for i := 1; i <= NumJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= NumJobs; i++ {
		fmt.Printf("Result : %d\n", <-res)
	}
}

func Worker(i int, jobs chan int, res chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", i, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", i, job)
		res <- job * 2
	}
}
