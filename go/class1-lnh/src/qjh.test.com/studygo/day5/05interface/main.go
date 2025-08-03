package main

import "fmt"

// 接口的实现
type animal interface {
	run()
	eat(string)
}

type cat struct {
	name string
	feet uint8
}
type chicken struct {
	alias string
	feet  uint8
}

func (c cat) run() {
	fmt.Println("走猫步")
}
func (c cat) eat(e string) {
	fmt.Printf("%s在吃%s\n", c.name, e)
}

func (c chicken) run() {
	fmt.Println("鸡动")
}
func (c chicken) eat(e string) {
	fmt.Printf("%s在吃%s\n", c.alias, e)
}

func main() {
	var a animal          // 定义一个接口类型的变量
	fmt.Printf("%T\n", a) //<nil>
	var a2 animal
	c1 := cat{ // 定义一个cat类型的变量bc
		name: "tom",
		feet: 4,
	}
	c2 := chicken{
		alias: "坤",
		feet:  2,
	}
	a = c1
	a.eat("猫粮")           //tom在吃猫粮
	fmt.Printf("%T\n", a) //main.cat
	fmt.Println(a)        //{tom 4}
	a2 = c2
	a2.eat("篮球") //坤在吃篮球
}
