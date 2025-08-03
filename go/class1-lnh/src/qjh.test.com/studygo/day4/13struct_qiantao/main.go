package main

import "fmt"

// 结构体嵌套
type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type person struct {
	name    string
	age     uint8
	address // 匿名嵌套结构体
	workPlace
	// address:address
}
type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "东尼大木",
		age:  40,
		address: address{
			province: "Japnan",
			city:     "Tykyo",
		},
		workPlace: workPlace{
			province: "Taiwan",
			city:     "Taipy",
		},
	}
	c1 := company{
		name: "Tykyo Hot",
		address: address{
			province: "California",
			city:     "Losangles",
		},
	}
	fmt.Println(p1.name, p1.address.city, p1.workPlace.city) //东尼大木 Tykyo Taipy
	// fmt.Println(p1.city) // 先在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(c1.name, c1.address.city, c1.address.province) //Tykyo Hot Losangles California
}
