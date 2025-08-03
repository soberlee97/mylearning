package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  uint8
}
type dog struct {
	name string
}

// 构造函数:约定成俗用new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针,减少程序的内存开销
func newPerson(name string, age uint8) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func newDog(dogName string) dog {
	return dog{
		name: dogName,
	}
}

func main() {
	p1 := newPerson("五月天", 114)
	fmt.Println(p1) //&{五月天 114}
	p2 := newDog("旺财")
	fmt.Println(p2) //{旺财}
}
