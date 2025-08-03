package main

import "fmt"

// 匿名字段
// 字段比较少也比较简单的场景
// 不常用!!!
type person struct {
	name string
	age  uint8
}

func main() {
	p1 := person{
		"劳斯莱斯",
		18,
	}
	fmt.Println(p1)              //{劳斯莱斯 18}
	fmt.Println(p1.age, p1.name) //18 劳斯莱斯
}
