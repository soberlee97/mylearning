package main

import (
	"fmt"
	"strings"
)

/*
	func main() {
		var string = "操你妈"
		if strings.HasSuffix(string, "逼") {
			fmt.Println("有逼")
		} else {
			fmt.Println("没逼")
		}
	}
*/
// 闭包
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("你好"))
	fmt.Println(jpgFunc("不好"))
	fmt.Println(jpgFunc("正常的.jpg"))
	fmt.Println(txtFunc("Maybach"))
	fmt.Println(txtFunc("Benz"))
}

/*
你好.jpg
不好.jpg
正常的.jpg
Maybach.txt
Benz.txt
*/
