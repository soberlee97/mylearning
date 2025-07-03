package main

import "fmt"

// pointer
func main() {
	// // 1. &:取地址
	//p := 18
	//n := &p
	//fmt.Printf("%T\n", n) //*int  int类型的指针
	// 2. *：根据地址取值
	//m := *n
	//fmt.Printf("%T,%d\n", m, m) //int,18
	var a1 *int
	fmt.Println(a1)   //<nil> pointer
	var a2 = new(int) // new函数申请一个内存地址  0x100a0fc
	fmt.Println(a2)
	fmt.Println(*a2) //0
	*a2 = 114514
	fmt.Println(*a2) //114514

}
