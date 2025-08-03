package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func menuList() {
	fmt.Println("---------欢迎来到学生管理系统")
	fmt.Println(`
	1. 查看所有学生
	2. 新增学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[uint64]student, 100),
	}
	for {
		menuList()
		var choice uint8
		fmt.Print("请输入序号:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			smr.listStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("滚~")
		}
	}
}
