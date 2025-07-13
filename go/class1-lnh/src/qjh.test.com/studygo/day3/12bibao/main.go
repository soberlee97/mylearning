package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	a := func(i int) int {
		base += i
		return base
	}
	b := func(i int) int {
		base -= i
		return base
	}
	return a, b
}
func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

}
