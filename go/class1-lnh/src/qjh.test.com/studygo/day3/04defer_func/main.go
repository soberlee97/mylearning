package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("好好好")
	defer fmt.Println("哟西")
	defer fmt.Println("end")
}
func main() {
	deferDemo()
}
