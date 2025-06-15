package main

import "fmt"

func main() {
	b1 := true
	fmt.Printf("%T\n", b1)
	var v2 bool // 默认是false
	fmt.Printf("%T, v2value=%v", v2, v2)
}
