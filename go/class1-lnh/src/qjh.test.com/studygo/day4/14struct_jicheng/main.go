package main

import "fmt"

// 结构体模拟实现其他语言中的"继承"
type animal struct {
	feet uint8
	name string
}

// 狗类
type dog struct {
	animal // animal拥有的方法,dog此时也有了
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s在移动\n", a.name)
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在大狗大狗叫叫叫\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			feet: 2,
			name: "大白",
		},
	}
	fmt.Println(d1) //{{2 大白}}
	d1.wang()       //大白在大狗大狗叫叫叫
	d1.move()       //大白在移动

}
