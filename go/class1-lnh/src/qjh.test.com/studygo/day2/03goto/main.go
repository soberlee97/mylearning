package main

import "fmt"

// goto
func main() {

	var flag = true

	for i := 0; i < 10; i++ {
		// 跳出多层for循环
		for j := 'A'; j < 'Z'; j++ {

			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			println("退出整体循环！")
			break // 跳出for循环（外层的for循环）
		}
	}
	// goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx // 跳到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
xx: // label标签
	fmt.Println("使用goto退出整体循环！")
}
