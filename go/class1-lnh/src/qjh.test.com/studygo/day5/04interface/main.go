package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车,都能跑
// 定义一个car接口类型
// 不管是什么结构体 只要有run方法都能是car类型
type car interface {
	run()
}

type benz struct {
	brand string
	speed uint16
}

type porsche struct {
	brand string
	speed uint16
}

func (b benz) run() {
	fmt.Printf("型号:%s,最高时速:%d\n", b.brand, b.speed)
}

func (p porsche) run() {
	fmt.Printf("型号:%s,最高时速:%d\n", p.brand, p.speed)
}

// drive函数接收一个car类型的变量
func drive(c car) {
	c.run()
}

func main() {
	maybache := benz{
		brand: "Maybache-s480",
		speed: 300,
	}
	plml911 := porsche{
		brand: "911",
		speed: 380,
	}

	drive(maybache) //型号:Maybache-s480,最高时速:300
	drive(plml911)  //型号:911,最高时速:380

}
