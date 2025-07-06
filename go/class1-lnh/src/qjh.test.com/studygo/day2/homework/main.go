package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1. 判断字符串中汉字的数量
// 难点是判断一个字符是汉字

func han_sum(m string) (ret int) {
	// // 1. 依次拿到字符串中的字符
	var count int
	for _, i := range m {
		// 2. 判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, i) {
			// 3. 把汉字出现的次数累加得到最终结果
			count++
		}
	}

	return count
}
func main() {
	fmt.Println(han_sum("ajhsdjka你好aaa"))

	//2. how do you do 单词出现的次数

	s1 := "how do you do"
	// 2.1 把字符串按照空格切割得到切片
	s2 := strings.Split(s1, " ")
	//fmt.Println(s2)

	// 2.2 遍历切片存储到一个map
	s3 := make(map[string]int, 10)
	for _, i := range s2 {
		if _, ok := s3[i]; !ok {
			s3[i] = 1
		} else {
			s3[i]++
		}
	}
	// 2.3 累加出现的次数
	fmt.Println(s3)
	for key, value := range s3 {
		fmt.Println(key, value)
	}

	// 回文判断
	// 字符串从左往右读和从右往左读是一样的，那么就是回文。
	// 上海自来水来自海上 s[0]  s[len(s)-1]
	// 山西运煤车煤运西山
	// 黄山落叶松叶落山黄

	ss := "a山西运煤车ss煤运西山a"
	//fmt.Println(len(ss))
	rune_ss := make([]rune, 0, len(ss))
	for _, v := range ss {
		rune_ss = append(rune_ss, v)
	}
	//fmt.Printf("%c", rune_ss) [a 山 西 运 煤 车 煤 运 西 山 a]
	for i := 0; i < len(rune_ss)/2; i++ {
		if rune_ss[i] != rune_ss[len(rune_ss)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
