package main

import "fmt"

// 空接口

// interface: 关键字
// interface{} :空接口类型

// 空接口作为函数参数,表示任意类型
func show(x interface{}) {
	fmt.Printf("%T,%v\n", x, x)
}

func main() {
	var any map[string]interface{}
	any = make(map[string]interface{}, 16)
	any["name"] = "老八"
	any["married"] = true
	any["age"] = 40
	any["hobby"] = []string{"吃屎", "rap", "直播", "做饭"}
	fmt.Println(any) //map[age:40 hobby:[吃屎 rap 直播 做饭] married:true name:老八]
	show(any)        //map[string]interface {},map[age:40 hobby:[吃屎 rap 直播 做饭] married:true name:老八]
	show(false)      //bool,false
	show(nil)        //<nil>,<nil>

}
