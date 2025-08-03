package main

import "fmt"

// 结构体

type person struct {
	name   string
	age    uint8
	gender string
	hoby   []string
}

func main() {
	// 声明一个person类型的变量p1
	var p1 person
	p1.name = "qjh"
	p1.age = 18
	p1.gender = "man"
	p1.hoby = []string{"唱", "跳", "rap", "篮球"}

	fmt.Printf("%T,%v\n", p1, p1) //main.person,{qjh 18 man [唱 跳 rap 篮球]}
	fmt.Println(p1)               //{qjh 18 man [唱 跳 rap 篮球]}
	// 访问变量p的字段
	fmt.Println(p1.name) //qjh

	var p2 person
	p2.name = "ymd"
	p2.age = 100
	fmt.Printf("%T,%v\n", p2, p2) //main.person,{ymd 100  []}

	// 匿名结构体:多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "tmp"
	s.y = 100
	fmt.Println(s)            //{tmp 100}
	fmt.Printf("%T,%v", s, s) //struct { x string; y int },{tmp 100}

}
