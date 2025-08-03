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

func (s *student) addStudent() {
	fmt.Print("请输入学生id:")
	fmt.Scanln(&s.id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&s.name)
	allStu[s.id] = s
}

func (s *student) listStudent() {
	for i, v := range allStu {
		fmt.Printf("学号:%v,姓名:%v\n", i, v.name)
	}
}
func (s *student) deleteStudent() {
	fmt.Printf("请输入需要删除的学生id:")
	var inputId uint64
	fmt.Scanln(&inputId)
	delete(allStu, inputId)
}
func main() {
	allStu = make(map[uint64]*student)
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
		l := newStudent(0, "tmp")
		fmt.Scanln(&inPut)
		switch inPut {
		case 1:
			l.listStudent()
		case 2:
			l.addStudent()
		case 3:
			l.deleteStudent()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("滚")
		}
	}
}
