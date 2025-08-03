package main

import (
	"fmt"

	tmp "day5/10calc"
)

func init() {
	fmt.Println("这是import_demo main导入时执行！")
}

func main() {
	a := tmp.Add(1, 2)
	fmt.Println(a)
}
