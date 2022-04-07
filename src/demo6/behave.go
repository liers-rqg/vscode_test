package demo6

import "fmt"

//接口
type ManBehave interface {
	Eat()
	Sleep()
	Play()
	Work()
}

type Children struct {
}
type Adult struct {
}

func (children *Children) Eat() {
	fmt.Println("children eating")
}
func (children *Children) Sleep() {
	fmt.Println("children sleeping")
}
func (children *Children) Play() {
	fmt.Println("children playing")
}
func (Children *Children) Work() {
	fmt.Println("children working")
}

func (adult *Adult) Eat() {
	fmt.Println("adult eating")
}
func (adult *Adult) Sleep() {
	fmt.Println("adult sleeping")
}
func (adult *Adult) Play() {
	fmt.Println("adult playing")
}
func (adult *Adult) Work() {
	fmt.Println("adult working")
}

func Recieve(manBehave ManBehave) {
	manBehave.Eat()
	manBehave.Sleep()
	manBehave.Play()
	manBehave.Work()
}
