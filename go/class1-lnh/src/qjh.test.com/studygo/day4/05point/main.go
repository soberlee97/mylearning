package main

import "fmt"

func main() {
	var a int = 100
	b := &a
	fmt.Printf("a:%T,b:%T\n", a, b) //a:int,b:*int
	// 将a的十六进制内存地址打印出来
	fmt.Printf("%p\n", &a) //0xc00000a0e8
	fmt.Printf("%p\n", b)  //0xc00000a0e8  // b的值
	fmt.Printf("%p\n", &b) //0xc000062060   // b的内存地址
}
