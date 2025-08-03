package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet uint8
}

// cat实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步")
}

// cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Println("吃", food)
}

func main() {
	var a1 animal
	var c1 = cat{
		name: "tom",
	}
	a1 = &c1
	a1.eat("大便") //吃 大便
	a1.move()    //走猫步
}
