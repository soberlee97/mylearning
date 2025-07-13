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



## 函数

### 常规

````go
package main

import "fmt"

// 函数:一段代码的封装

func f1() {
	fmt.Println("Hello 沙河！")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      // 如果使用命名的返回值,return后面可以省略返回值变量
}

// Go语言中支持多个返回值
func f7(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}


func main() {
	f1()
	f2("理想")
	f2("姬无命")
	fmt.Println(f3(100, 200)) // 调用函数

	ret := f3(100, 200)
	fmt.Println(ret)

	f5("lixiang", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// 在一个命名的函数中不能够再声明命名函数
	// func f8(){

	// }


}
````

#### day 3 homework

````go
package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	// 1. 依次拿到每个人的名字
	for _, name := range users {
		//distribution[name] = 0
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}

	}
	left = coins
	return

	// 2. 拿到一个人名根据分金币的规则去分金币,
	// 2.1 每个人分的金币数应该保存到 distribution 中
	// 2.2 还要记录下剩余的金币数
	// 3. 整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
	//return
}
func main() {
	//left := dispatchCoin()
	//fmt.Println("剩下：", left)
	left := dispatchCoin()
	for k, v := range distribution {
		fmt.Printf("%s,%d\n", k, v)
	}
	fmt.Println("剩下", left)
}
````



### 内置函数

|    内置函数    |                             介绍                             |
| :------------: | :----------------------------------------------------------: |
|     close      |                     主要用来关闭channel                      |
|      len       |      用来求长度，比如string、array、slice、map、channel      |
|      new       | 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针 |
|      make      |   用来分配内存，主要用来分配引用类型，比如chan、map、slice   |
|     append     |                 用来追加元素到数组、slice中                  |
| panic和recover |                        用来做错误处理                        |

### 作用域

#### 局部/全局 变量作用域

````go
package main

import "fmt"

// 变量的作用域
var x = 100 // 定义一个全局变量

// 定义一个函数
func f1() {
	//x = 114514
	name := "qjh"
	// 函数中查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 找不到就往函数的外面查找,一直找到全局
	fmt.Println(x, name)
}
func main() {
	f1()
    // fmt.Println(name) // 函数内部定义的变脸只能在该函数内部使用
}

````

#### 语句块作用域

````go
package main

import "fmt"


func main() {
	// 语句块作用域
	if i := 10; i < 18 {
		fmt.Println("乖乖上学")
	}
	// fmt.Println(i) // 不存在i
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j) // 不存在j
}
````

### 函数类型

````go
package main

import "fmt"

func f1() {
	fmt.Println("Hello 沙河！")
}

func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}
func ff(a, b int) int {
	return a + b
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}
func main() {
	a := f1
	fmt.Printf("%T\n", a) //func()
	b := f2
	fmt.Printf("%T\n", b)  //func() int
	f3(f2)                 //10
	f3(b)                  //10
	fmt.Printf("%T\n", f5) //func(func() int) func(int, int) int
	f7 := f5(f2)
	// f3(f4)
	fmt.Printf("%T", f7) //func(int, int) int

}
````



### 匿名函数

````go
package main

import "fmt"

// 匿名函数
/*var f1 = func(x, y int) {
	fmt.Println(x + y)

}*/

func main() {
	// 函数内部没有办法声明带名字的函数
	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(100, 200) //300

	// 如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("Hello world!")
	}(100, 200)
	//3000
	//Hello world!
}

````

### 高阶函数

函数也是一种类型，它可以作为参数，也可以作为返回值。

```go
// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}
```

### 闭包

````go
package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求：
// f1(f2)
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200) // 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	fmt.Printf("%T\n", ret)
	// ret()
	f1(ret)
}

//func()
//this is f1
//this is f2
//300
````

````go
package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

// 底层的原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找

func adder1() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x

	}
}

func main() {
	ret := adder(100)
	fmt.Println(ret(200)) //300
	fmt.Println(ret(300)) //600
	ret2 := ret(100)
	fmt.Println(ret2) //700

	f1 := adder(10)
	fmt.Println(f1(20)) //30
}
````

````go
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
````

#### 添加后缀示例

````go
package main

import (
	"fmt"
	"strings"
)

/*
	func main() {
		var string = "操你妈"
		if strings.HasSuffix(string, "逼") {
			fmt.Println("有逼")
		} else {
			fmt.Println("没逼")
		}
	}
*/
// 闭包
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("你好"))
	fmt.Println(jpgFunc("不好"))
	fmt.Println(jpgFunc("正常的.jpg"))
	fmt.Println(txtFunc("Maybach"))
	fmt.Println(txtFunc("Benz"))
}

/*
你好.jpg
不好.jpg
正常的.jpg
Maybach.txt
Benz.txt
*/
````

### defer

**💡 快速判断技巧**

遇到这类题时，问自己：

1. **返回值是否命名**？
   - 是 → defer 可直接修改返回值
   - 否 → 看 defer 是否通过闭包捕获了返回值（匿名返回值无法被直接捕获）
2. **defer 是否有参数**？
   - 有参数 → 操作的是副本，不影响外部
   - 无参数 → 闭包可能"偷家"改原件

**🚀 终极口诀**

> "**return 先赋值，defer 后捣乱，
> 闭包逮到就能改，传参只能玩副本**​"

#### defer 基础顺序流程

defer延迟调用，会把defer后面的语句延迟调用

把当时的状态都保存

defer多用于释放资源

多个defer存在时，按照先进后出的方式去执行。

````go
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

start
end
哟西
好好好
嘿嘿嘿
````



#### defer结合匿名函数

````go
package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x // 1. 返回值赋值 2. defer 3. 真正的RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x

}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x
	}()
	return x // 1. 返回值 = y = x = 5 2. defer修改的是x 3. 真正的返回
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (result int) {
	x := 5
	defer func() {
		x++
	}()
	result = x
	return // 返回result
}

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println(f5()) //5

}
````

#### defer会立即执行参数求值

````go
package main

import "fmt"

// defer

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //参数立即求值​
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 1. a:=1
// 2. b:=2
// 3. defer calc("1", 1, calc("10", 1, 2))
// 4. calc("10", 1, 2) // "10" 1 2 3
// 5. defer calc("1", 1, 3)
// 6. a = 0
// 7. defer calc("2", 0, calc("20", 0, 2))
// 8. calc("20", 0, 2) // "20" 0 2 2
// 9. defer calc("2", 0, 2)
// 10. b = 1
// calc("2", 0, 2) // "2" 0 2 2
// calc("1", 1, 3) // "1" 1 3 4

// 最终的答案：
// "10" 1 2 3
// "20" 0 2 2
//  "2" 0 2 2
// "1" 0 3 3
````

### panic和recover

用于Go的抛出错误与异常处理

使用`panic/recover`模式来处理错误。 `panic`可以在任何地方引发，但`recover`只有在`defer`调用的函数中有效。 首先来看一个例子：

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。

````go
package main

import "fmt"

func A() {
	fmt.Println("Plan A")
}
func B() {
	panic("出现致命错误！")
}
func C() {
	fmt.Println("Plan C")
}
func main() {
	A()
	B()
	C()
}

/*
Plan A
panic: 出现致命错误！

goroutine 1 [running]:
main.B(...)
        C:/Users/qjh/Desktop/运维/go/mylearning/mylearning/go/class1-lnh/src/qjh.test.com/studygo/day3/14panic_recover/main.go:9
main.main()
        C:/Users/qjh/Desktop/运维/go/mylearning/mylearning/go/class1-lnh/src/qjh.test.com/studygo/day3/14panic_recover/main.go:16 +0x5b
exit status 2
*/
````

使用recovery挽救

````go
package main

import (
	"fmt"
)

func A() {
	fmt.Println("Plan A")
}
func B() {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != nil {
			fmt.Println("已恢复PlanB")
		}
	}()
	panic("出现致命错误！")
}
func C() {
	fmt.Println("Plan C")
}
func main() {
	A()
	B()
	C()
}
/*
Plan A
出现致命错误！
已恢复PlanB
Plan C
*/
````

