package main

import "fmt"

type ff func()

// 类型断言1
func assign(x interface{}) {
	fmt.Printf("type:%T", x)
	str, ok := x.(string)
	if !ok {
		fmt.Println("断言失败，猜错了")
	} else {
		fmt.Println("猜对了，为string类型,值为:", str)
	}
}

// 类型断言2
func assign2(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Printf("这是一个数值型：%v\n", t)
	case string:
		fmt.Printf("这是一个字符串型:%v\n", t)
	case int64:
		fmt.Printf("这是一个int64型:%v\n", t)
	case []string:
		fmt.Printf("这是一个slice型:%v\n", t)
	case map[string]int:
		fmt.Printf("这是一个map型:%v\n", t)
	case bool:
		fmt.Printf("这是一个bool型:%v\n", t)
	case func():
		fmt.Println("这是一个函数型:", t)
	}
}
func f() {

}

func main() {

	assign(100)                                       //type:int断言失败，猜错了
	assign("逃避")                                      //type:string猜对了，为string类型,值为: 逃避
	assign2(10)                                       //这是一个数值型：10
	assign2("哈哈哈")                                    //这是一个字符串型:哈哈哈
	assign2([]string{"duck", "huck", "muck", "fuck"}) //这是一个slice型:[duck huck muck fuck]
	m1 := make(map[string]int, 16)
	m1["age"] = 18
	m1["height"] = 180
	m1["weight"] = 80
	assign2(m1)        //这是一个map型:map[age:18 height:180 weight:80]
	assign2(true)      //这是一个bool型:true
	assign2(int64(10)) //这是一个int64型:10
	assign2(f)         //这是一个函数型: 0x74c7e0
}
