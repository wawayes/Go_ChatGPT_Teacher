package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	jsonData := `
		{
			"name": "Tom",
			"age": 20,
			"address": "Beijing"
		}
	`
	var student Student
	err := json.Unmarshal([]byte(jsonData), &student)
	if err != nil {
		fmt.Println("解析出错：", err)
		return
	}
	fmt.Printf("姓名：%s，年龄：%d，地址：%s\n", student.Name, student.Age, student.Address)

	newStudent := Student{Name: "Jerry", Age: 25, Address: "Shanghai"}
	newJsonData, err := json.Marshal(newStudent)
	if err != nil {
		fmt.Println("序列化出错：", err)
		return
	}
	fmt.Printf("序列化结果：%s\n", string(newJsonData))
}
