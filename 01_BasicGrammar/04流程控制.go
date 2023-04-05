package main

import "fmt"

func demo004() {
	// if else 控制流程
	num := 10
	if num < 5 {
		fmt.Println("num is less than 5")
	} else if num < 10 {
		fmt.Println("num is between 5 and 10")
	} else {
		fmt.Println("num is greater than or equal to 10")
	}

	// for 循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while 循环
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 无限循环和 break
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	// switch 控制流程
	dayOfWeek := "Monday"
	switch dayOfWeek {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	case "Wednesday", "Thursday":
		fmt.Println("It's either Wednesday or Thursday!")
	default:
		fmt.Println("It's not Monday, Tuesday, Wednesday, or Thursday!")
	}
}
