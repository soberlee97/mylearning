package calc

import "fmt"

func init() {
	fmt.Println("导入calc包时执行")
}
func Add(x, y int) int {
	return x + y
}
