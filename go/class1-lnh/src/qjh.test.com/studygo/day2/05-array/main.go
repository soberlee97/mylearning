package main

import (
	"fmt"
)

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var (
		a1 [3]bool
		a2 [4]bool
	)
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)
	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""）
	fmt.Println(a1, a2) //[false false false] [false false false false]

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""）

	// 1. 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2. 初始化方式2：根据初始值自动推断数组的长度是多少
	var a10 = [...]int{1, 2, 3, 4, 5, 5, 6, 7}
	fmt.Println(len(a10), a10) //8 [1 2 3 4 5 5 6 7]
	// 3. 初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3) //[1 0 0 0 2]

	// 数组的遍历
	citys := [3]string{"无锡", "常州", "苏州"}
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	//[[1 2] [3 4] [7 8]]
	a11 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{7, 8},
	}
	fmt.Println(a11)
	// 多维数组的遍历
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组是值类型
	var b1 = [3]int{10, 20, 30}
	var b2 = b1
	b2[0] = 100
	fmt.Printf("%v,%v", b1, b2) //[10 20 30],[100 20 30]

}
