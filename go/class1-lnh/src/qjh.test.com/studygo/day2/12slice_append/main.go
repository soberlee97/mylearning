package main

import "fmt"

func main() {
	var a1 = [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]
	// 关于append删除切片中的某个元素
	// 删掉索引为1的那个3
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1) //[1 5 7 9 11 13 15 17]
	fmt.Println(a1) //[1 5 7 9 11 13 15 17 17]
}
