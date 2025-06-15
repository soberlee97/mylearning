package main

import "fmt"

/* 声明变量
var name string
var num int
var isOk bool
*/

//推荐使用小驼峰声明变量
var studentName string

// 批量声明，全局变量
var (
	name string
	num  int
	isOk bool
)

// s4 := "asd"

func main() {
	name = "理想迈巴赫"
	num = 114514
	isOk = true
	//Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("我的：%s\n", name) //%s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(num)            // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "hello-s1"
	fmt.Println(s1)

	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "hello-s2"
	fmt.Println(s2)

	// 简短变量声明，只能在函数里面用,不能在全局变量使用，常用
	s3 := "hello-s3"
	fmt.Println(s3)

	// s1 := "10" // 同一个作用域（{}）中不能重复声明同名的变量
	// 匿名变量是一个特殊的变量：_(后面学了函数再说)
}
