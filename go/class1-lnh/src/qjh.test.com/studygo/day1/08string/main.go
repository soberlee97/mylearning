package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "'C:\\Users\\HASEE\\Desktop\\learning-materials'"
	fmt.Println(path) //'C:\Users\HASEE\Desktop\learning-materials'
	s := "I'm ok"
	fmt.Println(s) //I'm ok

	//多行字符串
	s1 := `
	多行测试
	行一
	举头望明月
低头思故乡
								此致
							敬礼
	`
	fmt.Println(s1)

	s2 := `C:\Users\HASEE\Desktop\learning\go\class1-lnh\src\qjh.test.com\studygo\day1\04int`
	fmt.Println(s2)      //C:\Users\HASEE\Desktop\learning\go\class1-lnh\src\qjh.test.com\studygo\day1\04int
	fmt.Println(len(s2)) //字符串长度

	//字符串拼接
	dream := "理想"
	mayBach := "迈巴赫"
	ss := dream + mayBach
	//fmt.Printf("%s%s", dream, mayBach)
	fmt.Println(ss) //理想迈巴赫

	ret := strings.Split(s2, "\\")
	fmt.Println(ret) //[C: Users HASEE Desktop learning go class1-lnh src qjh.test.com studygo day1 04int]
	// 包含
	fmt.Println(strings.Contains(ss, "迈巴赫")) //true
	fmt.Println(strings.Contains(ss, "奇瑞"))  //false
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想")) //true  //判断包含
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "赫")) //true

	//获取index位置
	s3 := "akjashd"
	fmt.Println(strings.Index(s3, "a"))     //0
	fmt.Println(strings.LastIndex(s3, "a")) //3
	fmt.Println(strings.Index(s3, "av"))    //-1

	//拼接
	fmt.Println(strings.Join(ret, "+")) //C:+Users+HASEE+Desktop+learning+go+class1-lnh+src+qjh.test.com+studygo+day1+04int

}
