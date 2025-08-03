package main

import "fmt"

// 结构体占用一块连续的内存空间
type x struct {
	x int8 // 8bit => 1byte
	y int8
	z int8
}

func main() {
	m1 := x{
		x: int8(10),
		y: int8(20),
		z: int8(30),
	}
	fmt.Println(&m1.x) //0xc00000a0e8
	fmt.Println(&m1.y) //0xc00000a0e9
	fmt.Println(&m1.z) //0xc00000a0ea
}
