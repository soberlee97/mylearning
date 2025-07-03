package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            //copy(dest,src)
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 1000
	fmt.Println(a1, a2, a3) //[1000 3 5] [1000 3 5] [1 3 5]

	//将a1中的索引为1的3这个元素删掉
	//fmt.Printf("len(a1):%d,cap(a1):%d\n", len(a1), cap(a1)) //len(a1):3,cap(a1):3
	//a1 = append(a1[:1], a1[2:]...)
	//fmt.Printf("len(a1):%d,cap(a1):%d\n", len(a1), cap(a1)) //len(a1):2,cap(a1):3

	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	s1 := []int{1, 3, 5, 7, 9}
	x1 := s1
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(x1):%d cap(x1):%d\n", len(x1), cap(x1))
	fmt.Println(&s1[0], &x1[0])    //0x1014120 0x1014120
	x1 = append(s1[:1], s1[2:]...) // 修改了底层数组！
	fmt.Println(&s1[0], &x1[0])    //0x1014120 0x1014120   底层依然没变

	x1[0] = 100     // 修改底层数组
	fmt.Println(s1) //[100 5 7 9 9]

}
