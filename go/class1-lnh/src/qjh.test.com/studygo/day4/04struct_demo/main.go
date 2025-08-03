package main

import "fmt"

// 结构体是值类型
type person struct {
	name, gender string
}

// go语言中函数传参数永远传的是拷贝
func f1(n person) {
	n.gender = "女"
}

func f2(n *person) {
	//(*n).gender = "女" // 根据内存地址找到那个原变量,修改的就是原来的变量
	n.gender = "女" // 语法糖,自动根据指针找对应的变量
}

func main() {
	var p1 person
	p1.name = "qjh"
	p1.gender = "男"
	fmt.Println(p1)        //{qjh 男}
	fmt.Println(p1.gender) //男
	f1(p1)
	fmt.Println(p1.gender) //男
	f2(&p1)
	fmt.Println(p1.gender) //女

	var p2 = new(person)
	fmt.Printf("%T\n", p2) //*main.person
	*p2 = person{
		name:   "周玲",
		gender: "女",
	}
	fmt.Printf("%T %v\n", p2, p2) //*main.person &{周玲 女}

	var p3 = new(person)
	(*p3).name = "林心如"
	p3.gender = "女"         //语法糖
	fmt.Println(p3)         //&{林心如 女}
	fmt.Printf("%T\n", p3)  //*main.person
	fmt.Printf("%p\n", p3)  //0xc000066440  p3保存的值就是一个内存地址
	fmt.Printf("%p\n", &p3) //0xc000062068  求p3的内存地址
	// 2. 结构体指针2
	// 2.1 key-value初始化
	var p4 = &person{
		name: "哈哈哈哈",
	}
	fmt.Printf("%#v\n", p4) //&main.person{name:"哈哈哈哈", gender:""}
	// 2.2 使用值列表的形式初始化, 值的顺序要和结构体定义时字段的顺序一致
	var p5 = &person{
		"哈哈哈",
		"不男不女",
	}
	fmt.Printf("%#v\n", p5) //&main.person{name:"哈哈哈", gender:"不男不女"}

}
