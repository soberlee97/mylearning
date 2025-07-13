package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

// 底层的原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找

func adder1() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x

	}
}

func main() {
	ret := adder(100)
	fmt.Println(ret(200)) //300
	fmt.Println(ret(300)) //600
	ret2 := ret(100)
	fmt.Println(ret2) //700

	f1 := adder(10)
	fmt.Println(f1(20)) //30
}
