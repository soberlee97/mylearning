# Go笔记

[李文周博客](https://www.liwenzhou.com/posts/Go/golang-menu/)

## 一、环境搭建

### 1.1目录架构

![image-20250607145640993](typroimage/image-20250607145640993.png)

### 1.2 运行编译

####    go加速

````bash
go env -w GOPROXY=https://goproxy.cn,direct
go env
````

####    go编译

````bash
go build hello.go
#生成exe可执行文件

go build -o hello-rename.exe hello.go
#重命名exe文件
````

####     go直接执行

````bash
go run hello.go
````

#### go跨平台编译(交叉编译)

````bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build
````



#### 1.3 Hello World!

````GO
package main

// 导入语句
import "fmt"

// 函数外只能防止标识符（变量\常量\函数\类型）的声明
// fmt.Println("Hello") // 非法

// 程序的入口函数
func main() {
	fmt.Println("Hello world!")
}
````

**main函数是程序的入口，package如果是main，go build出来的是个二进制文件，而且必须存在main函数，为程序入口的函数。**

## 变量和常量

### 变量声明，变量赋值

变量声明格式

步骤： 申明->赋值->使用

`var 变量名 变量类型`

````go
package main

import "fmt"

/* 声明变量
var name string
var num int
var isOk bool
*/

//推荐使用小驼峰声明变量
var studentName string

// 批量声明，全局变量
var (
	name string
	num  int
	isOk bool
)

// s4 := "asd"

func main() {
	name = "理想迈巴赫"
	num = 114514
	isOk = true
	//Go语言中非全局变量声明后必须使用，不使用就编译不过去
	fmt.Print(isOk) // 在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("我的：%s\n", name) //%s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(num)            // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "hello-s1"
	fmt.Println(s1)

	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "hello-s2"
	fmt.Println(s2)

	// 简短变量声明，只能在函数里面用,不能在全局变量使用，常用
	s3 := "hello-s3"
	fmt.Println(s3)

	// s1 := "10" // 同一个作用域（{}）中不能重复声明同名的变量
	// 匿名变量是一个特殊的变量：_(后面学了函数再说)
}


//输出结果
true
我的：理想迈巴赫
114514
hello-s1
hello-s2
hello-s3

````



### 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值。

**iota** 是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中**每新增一行常量声明将使`iota`计数一次**(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

````go
package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415

// 批量声明常量
const (
	codeOk   = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	_         //1   匿名变量丢弃
	b2        //2
	b3        //3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1=1 d2=2
	d3, d4 = iota + 1, iota + 2 // d3=2 d4=3
)

// 定义数量级
//定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
//也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，
//也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota) //2^10 1024
	MB = 1 << (10 * iota) //2^20 1048576
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//	fmt.Println("pi:", pi)            //pi: 3.1415
	//	fmt.Println("codeOk:", codeOk)    //codeOk: 200
	//	fmt.Println("notFound", notFound) //notFound 404
	//	fmt.Println("n1:", n1)            //n1: 100
	//	fmt.Println("n2:", n2)            //n2: 100
	//	fmt.Println("n3:", n3)            //n3: 100
	//fmt.Println("a1:", a1)
	//fmt.Println("a2:", a2)
	//fmt.Println("a3:", a3)
	//fmt.Println("b1:", b1)
	//fmt.Println("b2:", b2)
	//fmt.Println("b3:", b3)
	//fmt.Println("c1:", c1)
	//fmt.Println("c2:", c2)
	//fmt.Println("c3:", c3)
	//fmt.Println("c4:", c4)
	//fmt.Println("d1:", d1)
	//fmt.Println("d2:", d2)
	//fmt.Println("d3:", d3)
	//fmt.Println("d4:", d4)
	fmt.Println("1KB=", KB, "B")
	fmt.Println("1MB=", MB, "B")
}
````

### 整型

````go
package main

import "fmt"

func main() {
	var i1 = 114514
	fmt.Printf("%d\n", i1) //十进制
	fmt.Printf("%b\n", i1) //二进制
	fmt.Printf("%o\n", i1) //八进制
	fmt.Printf("%x\n", i1) //十六进制

	i2 := 077 //赋值八进制
	fmt.Printf("i2转化10进制: %d\n", i2)

	i3 := 0x12345 //赋值十六进制
	fmt.Printf("i3转化10进制: %d\n", i3)
	// 声明int8类型的变量
	i4 := int8(11)       //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T", i4) //int8
}
````

### 浮点型

````go
package main

import (
	"fmt"
)

// 浮点数
func main() {
	//math.MaxFloat64 // float32最大值
	f1 := 3.1415926
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型

	f2 := float32(3.1415926) // 显示声明float32类型
	fmt.Printf("%T", f2)
	// f1 = f2   // float32类型的值不能直接复赋值给float64类型的变量
	//f1 = float64(f2)
}
````

### 布尔型

````go
package main

import "fmt"

func main() {
	b1 := true
	fmt.Printf("%T\n", b1)
	var v2 bool // 默认是false
	fmt.Printf("%T, v2value=%v", v2, v2)
}
````

### fmt格式占位符,类型

````go
package main

import "fmt"

func main() {
	var n = 100
	fmt.Printf("%T\n", n)  //打印值的类型
	fmt.Printf("%d\n", n)  //十进制
	fmt.Printf("%b\n", n)  //二进制
	fmt.Printf("%o\n", n)  //八进制
	fmt.Printf("%x\n", n)  //十六进制
	fmt.Printf("%v\n", n)  //值的默认格式表示
	var s = "hello！我是艾卡大王"
	fmt.Printf("字符串:%s\n", s)  //字符串打印
	fmt.Printf("字符串:%v\n", s)  //字符串:hello！我是艾卡大王  
	fmt.Printf("字符串:%#v", s)   //字符串:"hello！我是艾卡大王"   //会加引号

}
````

### 字符串

#### 字符串常用方法

|                方法                 |      介绍      |
| :---------------------------------: | :------------: |
|              len(str)               |     求长度     |
|           +或fmt.Sprintf            |   拼接字符串   |
|            strings.Split            |      分割      |
|          strings.contains           |  判断是否包含  |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) |    join操作    |

````go
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
````

### rune型与byte型

````go
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
````

### if

````go
package main

import "fmt"

func main() {
	// age := 19

	// if条件判断
	/*	if age >= 18 {
			fmt.Println("澳门皇冠线上赌场开业啦！")
		} else {
			fmt.Println("快写暑假作业吧！")
		}
	*/

	// 多个判断条件
	/*	if age >= 35 {
			fmt.Println("赶紧去上班吧")
		} else if age >= 18 {
			fmt.Println("赶紧毕业吧！")
		} else {
			fmt.Println("快写暑假作业！")
		}
	*/

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("你已经成年啦！")
	} else {
		fmt.Println("快写暑假作业吧！")
	}
	age := 114514
	fmt.Println(age) //114514

}
````

### for循环

````go
package main

import (
	"fmt"
)

// for循环
func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1
	/* var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) //10
	*/

	// 变种2
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "Hello老钱"
	for i, v := range s {
		fmt.Printf("%d,%c\n", i, v)
	}

}
````

### 九九乘法表

````go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()

	}
}
````

