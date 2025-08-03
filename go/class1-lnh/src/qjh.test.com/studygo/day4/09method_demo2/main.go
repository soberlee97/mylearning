package main

import "fmt"

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法,只能给自己包里的类型添加方法
type myInt int

func (m myInt) Hello() {
	fmt.Printf("这是一个int，值为%d", m)
}

func main() {
	m1 := myInt(114514)
	m1.Hello()
}
