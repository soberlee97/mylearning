package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1)              //map[]
	fmt.Println(m1 == nil)       //true    // 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 5) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["理想"] = 680
	m1["姬无命"] = 35

	fmt.Println(m1)        //map[姬无命:35 理想:680]
	fmt.Println(m1["姬无命"]) //35
	fmt.Println(m1["娜扎"])  //0
	// 约定成俗用ok接收返回的布尔值
	value, ok := m1["娜扎"] // 如果不存在这个key拿到对应值类型的零值
	if !ok {
		fmt.Println("查无此key！")
	} else {
		fmt.Printf("%v", value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "姬无命")
	fmt.Println(m1)
	delete(m1, "ssss") // 删除不存在的key)

}
