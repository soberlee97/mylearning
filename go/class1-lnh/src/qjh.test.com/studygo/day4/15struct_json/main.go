package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1.序列化:   把Go语言中的结构体变量 --> json格式的字符串
// 2.反序列化: json格式的字符串   --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	p1 := person{
		Name: "大大",
		Age:  18,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("无法完成序列化,错误为：%v", err)
		return
	}
	fmt.Println(b)                //[123 34 110 97 109 101 34 58 34 229 164 167 229 164 167 34 44 34 97 103 101 34 58 49 56 125]
	fmt.Printf("%v\n", string(b)) //{"name":"大大","age":18}
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)          //main.person{Name:"理想", Age:0x12}

}
