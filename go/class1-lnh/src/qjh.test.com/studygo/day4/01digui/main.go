package main

import "fmt"

// 递归:函数自己调用自己!
// 递归适合处理那种问题相同\问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 3! = 3*2*1     = 3*2!
// 4! = 4*3*2*1   = 4*3!
// 5! = 5*4*3*2*1 = 5*4!

// 计算n的阶乘
func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

// 上台阶的面试题
// n个台阶，一次可以走1步，也可以走2步，有多少种走法。
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
		// 如果只有一个台阶就一种走法
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	//ret := f1(7)
	//fmt.Println(ret)
	rettj := taijie(3)
	fmt.Println(rettj)
}
