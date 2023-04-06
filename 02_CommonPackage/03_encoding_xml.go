package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name    string `xml:"name"`
	Age     int    `xml:"age"`
	Address string `xml:"address"`
}

func demo003() {
	xmlData := `
		<person>
			<name>Tom</name>
			<age>12</age>
			<address>TaiYuan</address>
		</person>
	`
	var person Person
	err := xml.Unmarshal([]byte(xmlData), &person)
	if err != nil {
		return
	}

	fmt.Printf("将xmlData解析为结构体: Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)

	p := Person{Name: "Walker", Age: 22, Address: "Taiyuan"}
	marshal, err := xml.Marshal(p)
	if err != nil {
		return
	}
	fmt.Printf("将person结构体序列化为xml: %s", string(marshal))

}
