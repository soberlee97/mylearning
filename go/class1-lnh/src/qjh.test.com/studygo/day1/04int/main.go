package main

import "fmt"

func main() {
	var i1 = 114514
	fmt.Printf("%d\n", i1) //十进制
	fmt.Printf("%b\n", i1) //二进制
	fmt.Printf("%o\n", i1) //八进制
	fmt.Printf("%x\n", i1) //十六进制

	i2 := 077 //赋值八进制
	fmt.Printf("i2转化10进制: %d\n", i2)

	i3 := 0x12345 //赋值十六进制
	fmt.Printf("i3转化10进制: %d\n", i3)
	// 声明int8类型的变量
	i4 := int8(11)       //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T", i4) //int8
}
