package main

import "fmt"

func main() {
	s := "Hello沙河사샤"
	//n := len(s)
	//for i := 0; i < n; i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n", s[i])
	//}
	for _, c := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", c) // %c:字符
	}
	// "Hello" => 'H' 'e' 'l' 'l' 'o'

	//实现字符串替换
	s2 := "白萝卜"      //字符串不能修改   // => '白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	//fmt.Printf("%v,%c", s3[0], s3[0])
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	c1 := "红" //string
	c2 := '红' //int32
	fmt.Printf("c1:%T,c2:%T\n", c1, c2)

	c3 := "H"       //string
	c4 := byte('H') //uint8
	fmt.Printf("c3:%T,c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4) //72

	//类型转换
	n1 := 10 //int
	var f float64 = float64(n1)
	fmt.Println(f)        //10
	fmt.Printf("%T\n", f) //float64
}
