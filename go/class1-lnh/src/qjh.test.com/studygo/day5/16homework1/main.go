package main

import (
	"fmt"
	"os"
)

type myLog struct {
	//level string
	outPutType     string
	outPutFileName string
}

func NewmyLog(outPutType string, outPutFileName string) myLog {
	return myLog{
		//level: level,
		outPutType:     outPutType,
		outPutFileName: outPutFileName,
	}
}
func (l myLog) Trace(outPutString string) {
	switch l.outPutType {
	case "file_log":
		fileObj, err := os.OpenFile(l.outPutFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("openfile FAILED,err:", err)
			return
		}
		defer fileObj.Close()
		fmt.Fprintf(fileObj, "Trace log:%s\n", outPutString)
	case "stdout":
		fmt.Fprintf(os.Stdout, "Trace log:%s", outPutString)
	default:
		fmt.Println("outPutType参数只有\"file_log\"和\"stdout\"")
		os.Exit(0)
	}

}
func (l myLog) Debug(outPutString string) {
	switch l.outPutType {
	case "file_log":
		fileObj, err := os.OpenFile(l.outPutFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("openfile FAILED,err:", err)
			return
		}
		defer fileObj.Close()
		fmt.Fprintf(fileObj, "Debug log:%s\n", outPutString)
	case "stdout":
		fmt.Fprintf(os.Stdout, "Debug log:%s", outPutString)
	default:
		fmt.Println("outPutType参数只有\"file_log\"和\"stdout\"")
		os.Exit(0)
	}

}
func (l myLog) Warning(outPutString string) {
	switch l.outPutType {
	case "file_log":
		fileObj, err := os.OpenFile(l.outPutFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("openfile FAILED,err:", err)
			return
		}
		defer fileObj.Close()
		fmt.Fprintf(fileObj, "Warning log:%s\n", outPutString)
	case "stdout":
		fmt.Fprintf(os.Stdout, "Warning log:%s", outPutString)
	default:
		fmt.Println("outPutType参数只有\"file_log\"和\"stdout\"")
		os.Exit(0)
	}

}
func (l myLog) Info(outPutString string) {
	switch l.outPutType {
	case "file_log":
		fileObj, err := os.OpenFile(l.outPutFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("openfile FAILED,err:", err)
			return
		}
		defer fileObj.Close()
		fmt.Fprintf(fileObj, "Info log:%s\n", outPutString)
	case "stdout":
		fmt.Fprintf(os.Stdout, "Info log:%s", outPutString)
	default:
		fmt.Println("outPutType参数只有\"file_log\"和\"stdout\"")
		os.Exit(0)
	}

}
func (l myLog) Error(outPutString string) {
	switch l.outPutType {
	case "file_log":
		fileObj, err := os.OpenFile(l.outPutFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("openfile FAILED,err:", err)
			return
		}
		defer fileObj.Close()
		fmt.Fprintf(fileObj, "Error log:%s\n", outPutString)
	case "stdout":
		fmt.Fprintf(os.Stdout, "Error log:%s", outPutString)
	default:
		fmt.Println("outPutType参数只有\"file_log\"和\"stdout\"")
		os.Exit(0)
	}

}
func main() {
	l1 := myLog{
		outPutType:     "file_log",
		outPutFileName: "./myTest.log",
	}
	l1.Trace("这是日志测试2")
	l2 := myLog{
		outPutType: "stdout",
	}
	l2.Trace("这是终端输出！")
	l1.Debug("这是一个debug操作")
	l1.Info("这是info日志！")
	l1.Warning("这是警告日志！")
	l1.Error("这是错误日志！")

}
