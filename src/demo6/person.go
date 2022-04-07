package demo6

import "fmt"

//封装案例
type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//setter getter 方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 100 {
		p.age = age
	} else {
		fmt.Println("age error")
	}
}

func (p *person) SetSalary(salary float64) {
	if salary > 0 {
		p.salary = salary
	} else {
		fmt.Println("salary error")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

func (p *person) GetAge() int {
	return p.age
}
