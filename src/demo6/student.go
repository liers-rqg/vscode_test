package demo6

import "strconv"

//继承案例
type Student struct {
	Name  string
	age   int
	score float64
}

func (stu *Student) SetAge(age int) {
	stu.age = age
}
func (stu *Student) SetScore(score float64) {
	stu.score = score
}
func (stu *Student) ShowInfo() string {
	return "Student name = " + stu.Name + " age = " + strconv.Itoa(stu.age) + " score = " + strconv.FormatFloat(stu.score, 'f', 1, 64)
}

type Pupil struct {
	Student
}

type Granduate struct {
	Student
}

// type phd struct {
// 	//结构体里面可以嵌套基本数据类型
// 	Student
// 	//int
// }
