package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格
func useScan() {
	var s string
	fmt.Print("请输入:")
	fmt.Scanln(&s)
	fmt.Printf("输出内容为:%s\n", s)
}
func useBufio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入:")
	s, _ := reader.ReadString('\n')
	fmt.Printf("输入内容为：%s\n", s)

}

func main() {
	//useScan()  //无法空格输入输出
	useBufio()
	// logger.Trace()
	// logger.Debug()
	// logger.Warning()
	// logger.Info()
	// logger.Error("日志的内容")

	// 写日志
	fmt.Fprintln(os.Stdout, "这是一条输出！")
	fileObj, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	fmt.Fprintln(fileObj, "后续写入1")

}
