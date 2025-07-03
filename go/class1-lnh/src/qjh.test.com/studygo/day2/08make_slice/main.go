package main

import "fmt"

func main() {
	// make()函数创造切片
	s1 := make([]int, 5, 10)
	s2 := make([]int, 0, 10)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2):%d cap(s2):%d\n", s2, len(s2), cap(s2))

	// 切片的赋值
	s3 := []int{1, 3, 5, 7, 9}
	s4 := s3            // s3和s4都指向了同一个底层数组
	fmt.Println(s3, s4) //[1 3 5 7 9] [1 3 5 7 9]
	s3[0] = 100
	fmt.Println(s3, s4) //[100 3 5 7 9] [100 3 5 7 9]

	// 切片的遍历
	// 1. 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	// 2. for range循环
	for _, i := range s3 {
		fmt.Printf("%d\n", i)
	}

}
