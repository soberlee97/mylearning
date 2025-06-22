package main

import "fmt"

func main() {
	// age := 19

	// if条件判断
	/*	if age >= 18 {
			fmt.Println("澳门皇冠线上赌场开业啦！")
		} else {
			fmt.Println("快写暑假作业吧！")
		}
	*/

	// 多个判断条件
	/*	if age >= 35 {
			fmt.Println("赶紧去上班吧")
		} else if age >= 18 {
			fmt.Println("赶紧毕业吧！")
		} else {
			fmt.Println("快写暑假作业！")
		}
	*/

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("你已经成年啦！")
	} else {
		fmt.Println("快写暑假作业吧！")
	}
	age := 114514
	fmt.Println(age) //114514

}
