package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型

type myInt int64   // 自定义类型
type yourInt = int // 类型别名

func main() {
	var a myInt = 200
	fmt.Printf("%T\n", a) //main.myInt
	b := yourInt(100)
	fmt.Printf("%T\n", b) //int

	var c rune
	c = '好'
	fmt.Printf("%T\n", c) //int32

}
