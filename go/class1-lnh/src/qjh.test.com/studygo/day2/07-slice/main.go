package main

import "fmt"

func main() {
	var q1 = []int{}    //初始化了
	var q2 = []string{} //初始化了

	fmt.Println(q1 == nil) //false
	fmt.Println(q2 == nil) //false

	var s1 []int
	var s2 []string
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"无锡", "常州", "苏州"}
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false

	//长度与容量
	fmt.Printf("len(s1):%v,cap(s1):%v\nlen(s2):%v,cap(s2):%v\n", len(s1), cap(s1), len(s2), cap(s2))
	//len(s1):3,cap(s1):3
	//len(s2):3,cap(s2):3

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含，（左闭右开）   [1 3 5 7]
	s4 := a1[1:6]
	fmt.Println(s3, s4)
	s5 := a1[:3] //[0:3]=>[1 3 5]
	s6 := a1[4:] //[4:len(a1)]  [9 11 13]
	s7 := a1[:]  // => [0:len(a1)]    [3 5 7 9 11 13]
	fmt.Println(s5, s6, s7)
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d,cap(s5):%d\n", len(s5), cap(s5)) //len(s5):3,cap(s5):7
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d,cap(s6):%d\n", len(s6), cap(s6)) //len(s6):3,cap(s6):3

	// 切片再切割
	s8 := s6[2:]
	fmt.Println(s8)                                         //[13]
	fmt.Printf("len(s8):%d,cap(s8):%d\n", len(s8), cap(s8)) //len(s8):1,cap(s8):1

	a1[6] = 1000           // 修改了底层数组的值,底层数组变了，切片里的也会变
	fmt.Println("s6:", s6) //s6: [9 11 1000]
	fmt.Println("s8:", s8) //s8: [1000]

}
