package main

import "fmt"

// 方法

// 标识符:变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的,就表示对外部包可见(暴露的,公有的).

// dog 这是一个狗的结构体

type person struct {
	name string
	age  int
}
type dog struct {
	name string
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量,多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s在大狗大狗叫叫叫\n", d.name)
}

// 使用值接收者:传拷贝进去
// func (p person) guonian() {
// 	p.age++
// }
func (p person) guonian() {
	p.age++
}

// 指针接收者:传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

func main() {
	d1 := newDog("大白")
	d1.wang() //大白在大狗大狗叫叫叫
	p1 := newPerson("jacky", 18)
	p1.guonian()
	//fmt.Println(p1.age)
	p2 := newPerson("lucy", 18)
	p2.zhenguonian()
	fmt.Println(p1.age, p2.age) //18 19

}
