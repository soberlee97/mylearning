package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 打开文件
func readFromFile1(s string) {
	fileObj, err := os.Open(s)
	if err != nil {
		fmt.Println("open file failed,error:", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128) // 指定读的长度
	for {
		var m [128]byte
		n, err := fileObj.Read(m[:])
		if err == io.EOF {
			fmt.Println("读完啦")
			return
		}
		if err != nil {
			fmt.Println("read from file failed,err:", err)
			return
		}
		fmt.Printf("读取了%d的字节数", n)
		fmt.Println(string(m[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用bufio这个包读取文件
func readFromFileBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("OpenFile failed,error:", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完啦")
			return
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Println(string(line))
	}

}

func readFileFromOs() {
	fileObj, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println("readFile failed,err:", err)
		return
	}
	fmt.Println(string(fileObj))
}

func main() {
	//readFromFile1("./main.go")
	//readFromFileBufio()
	readFileFromOs()

}
