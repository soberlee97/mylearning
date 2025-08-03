package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件写内容
// 0100 0000
func writeFileDemo1() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("openFileField,err:", err)
		return
	}
	fileObj.Write([]byte("你好\n"))
	fileObj.WriteString("我是第二行！\n")
	fileObj.Close()
}

func writeFileFromBufio() {
	fileObj, err := os.OpenFile("222.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("openFile failed,err:", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	write := bufio.NewWriter(fileObj)
	write.WriteString("如果我是DJ你会爱我吗！") // 写到缓存中
	write.Flush()                     // 将缓存中的内容写入文件
}

func writeFileFromOs(s string) {
	err := os.WriteFile("./os.txt", []byte(s), 0644)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}

}
func main() {
	writeFileDemo1()
	writeFileFromBufio()
	writeFileFromOs("os测试！")
}
