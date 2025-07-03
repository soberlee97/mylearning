package main

import "fmt"

// map和slice组合

func main() {
	// 元素类型为map的切片
	s1 := make([]map[int]string, 10, 10)
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "北京"
	fmt.Println(s1)
	// 值为切片类型的map
	var m1 = make(map[string][]string, 10)
	m1["北京"] = []string{"中关村", "饭岛爱", "麻生希"}
	fmt.Println(m1)
}
