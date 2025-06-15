package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415

// 批量声明常量
const (
	codeOk   = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	_         //1   匿名变量丢弃
	b2        //2
	b3        //3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1=1 d2=2
	d3, d4 = iota + 1, iota + 2 // d3=2 d4=3
)

// 定义数量级
//定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
//也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，
//也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota) //2^10 1024
	MB = 1 << (10 * iota) //2^20 1048576
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//	fmt.Println("pi:", pi)            //pi: 3.1415
	//	fmt.Println("codeOk:", codeOk)    //codeOk: 200
	//	fmt.Println("notFound", notFound) //notFound 404
	//	fmt.Println("n1:", n1)            //n1: 100
	//	fmt.Println("n2:", n2)            //n2: 100
	//	fmt.Println("n3:", n3)            //n3: 100
	//fmt.Println("a1:", a1)
	//fmt.Println("a2:", a2)
	//fmt.Println("a3:", a3)
	//fmt.Println("b1:", b1)
	//fmt.Println("b2:", b2)
	//fmt.Println("b3:", b3)
	//fmt.Println("c1:", c1)
	//fmt.Println("c2:", c2)
	//fmt.Println("c3:", c3)
	//fmt.Println("c4:", c4)
	//fmt.Println("d1:", d1)
	//fmt.Println("d2:", d2)
	//fmt.Println("d3:", d3)
	//fmt.Println("d4:", d4)
	fmt.Println("1KB=", KB, "B")
	fmt.Println("1MB=", MB, "B")
}
