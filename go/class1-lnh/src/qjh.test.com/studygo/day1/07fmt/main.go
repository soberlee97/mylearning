package main

import "fmt"

func main() {
	var n = 100
	fmt.Printf("%T\n", n) //打印值的类型
	fmt.Printf("%d\n", n) //十进制
	fmt.Printf("%b\n", n) //二进制
	fmt.Printf("%o\n", n) //八进制
	fmt.Printf("%x\n", n) //十六进制
	fmt.Printf("%v\n", n) //值的默认格式表示
	var s = "hello！我是艾卡大王"
	fmt.Printf("字符串:%s\n", s) //字符串打印
	fmt.Printf("字符串:%v\n", s) //字符串:hello！我是艾卡大王
	fmt.Printf("字符串:%#v", s)  //字符串:"hello！我是艾卡大王"   //会加引号

}
