package main

import "fmt"

const (
	eat   = 4 //100
	sleep = 2 //010
	beat  = 1 //001
)

func bin(arg int) {
	fmt.Printf("%b\n", arg)
}
func main() {
	bin(eat | beat)         //101
	bin(eat | sleep)        //110
	bin(eat | beat | sleep) //111
}
