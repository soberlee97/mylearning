package main

import (
	"fmt"
	"sort"
)

// 切片的练习题
func main() {
	var a1 = make([]int, 5, 10) // 创建切片，长度为5，容量为10
	for i := 0; i < 10; i++ {
		a1 = append(a1, i)
	}
	fmt.Println(a1, cap(a1)) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9] 20

	var a2 = [...]int{3, 8, 1, 0, 7, 11, 4}
	sort.Ints(a2[:]) // 对切片进行排序
	fmt.Println(a2)  //[0 1 3 4 7 8 11]

}
