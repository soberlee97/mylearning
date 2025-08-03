package main

import "fmt"

type student struct {
	id   uint64
	name string
}

type studentMgr struct {
	allStudent map[uint64]student
}

func (s studentMgr) listStudents() {
	for _, v := range s.allStudent {
		fmt.Printf("学号:%d,学生姓名:%s\n", v.id, v.name)
	}
}

func (s studentMgr) addStudent() {
	var (
		inputID   uint64
		inputName string
	)
	fmt.Print("请输入需要添加的学生id:")
	fmt.Scanln(&inputID)
	fmt.Print("请输入需要添加的学生姓名:")
	fmt.Scanln(&inputName)
	newStu := student{
		id:   inputID,
		name: inputName,
	}
	s.allStudent[inputID] = newStu
	fmt.Printf("%v添加成功！", s.allStudent[inputID])
}

func (s studentMgr) editStudent() {
	var (
		inputID   uint64
		inputName string
	)
	fmt.Print("请输入需要修改的学生id:")
	fmt.Scanln(&inputID)
	stuObj, ok := s.allStudent[inputID]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Println("您要修改的学生是:", stuObj)
	fmt.Print("请修改对应学生的姓名:")
	fmt.Scanln(&inputName)
	stuObj.name = inputName
	s.allStudent[inputID] = stuObj
}

func (s studentMgr) deleteStudent() {
	var (
		stuID uint64
	)
	fmt.Print("请输入要删除的学生id:")
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人请重试!")
		return
	}
	stuObj = s.allStudent[stuID]
	delete(s.allStudent, stuObj.id)
}
