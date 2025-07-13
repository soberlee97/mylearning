package main

import (
	"fmt"
)

func A() {
	fmt.Println("Plan A")
}
func B() {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != nil {
			fmt.Println("已恢复PlanB")
		}
	}()
	panic("出现致命错误！")
}
func C() {
	fmt.Println("Plan C")
}
func main() {
	A()
	B()
	C()
}
