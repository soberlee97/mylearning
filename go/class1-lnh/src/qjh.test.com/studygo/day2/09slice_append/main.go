package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	//s1[3] = "西雅图"  // 错误的写法 会导致编译错误：索引越界
	//fmt.Println(s1)

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	// 必须用变量接收append的返回值
	s1 = append(s1, "成都")
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "广州", "武汉", "辽宁")
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))

	tmp_city := []string{"常州", "无锡", "苏州", "镇江"}
	s1 = append(s1, tmp_city...) // ...表示拆开
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
}
