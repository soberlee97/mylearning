package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret

}
func main() {
	a := 2
	b := 3
	defer calc("11", a, calc("10", a, b)) //参数立即求值​
	a = 5
	defer calc("12", a, calc("22", a, b))
	b = 6
}

//1. a:=2
//2. b:=3
//3. defer calc("11",2,calc("10",2,3))
//4. calc("10",2,3) // "10" 2 3  5
//5. defer calc("11",2,5)
//6. a=5
//7. defer calc("12", a, calc("22", 5, 3))
//8. calc("22", 5, 3) // "22" 5 3 8
//9. defer calc("12",5,8)
//10. b=6
//推测为：
//"10" 2 3  5
//"22" 5 3 8
//"12" 5 8 13
//"11" 2 5 7
