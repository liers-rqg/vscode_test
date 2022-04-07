package demo5

import "strconv"

//工厂模式
type student struct {
	Name   string
	Grade  int
	Number string
}

//student属于私有结构体
func NewStudent(name string, grade int, number string) *student {
	return &student{
		Name:   name,
		Grade:  grade,
		Number: number,
	}
}

//String方法可以被print方法识别
func (stu *student) String() string {
	return "student name = " + stu.Name + " Grade = " + strconv.Itoa(
		stu.Grade,
	) + " Number = " +
		stu.Number
}
