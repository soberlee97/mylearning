package main

import (
	"fmt"
	"os"
)

type student struct {
	id   uint64
	name string
}

func newStudent(id uint64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

var (
	allStu map[uint64]*student
)

func addStudent() {
	var (
		id   uint64
		name string
	)
	fmt.Print("请输入学生id!\n")
	fmt.Scanln(&id)
	fmt.Printf("请输入学生姓名!\n")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStu[id] = stu
}

func listAll() {
	for i, v := range allStu {
		fmt.Printf("学号:%v,学生姓名:%v", i, v.name)
	}
}

func deleteStu() {
	var id uint64
	fmt.Print("请输入所需要删除的学生id:")
	fmt.Scanln(&id)
	found := false
	for i, _ := range allStu {
		if id == i {
			delete(allStu, id)
		}
		found = true
	}
	if !found {
		fmt.Println("该学生不存在请重新输入！")
	}
}

func main() {
	allStu = make(map[uint64]*student, 48)
	for {
		fmt.Println("欢迎来带学生管理系统！")
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出系统
	`)
		fmt.Print("请输入选项:")
		var inPut uint8
		fmt.Scanln(&inPut)
		switch inPut {
		case 1:
			listAll()
		case 2:
			addStudent()
		case 3:
			deleteStu()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("滚！")

		}

	}
}
