package demo3

//结构体
import (
	"fmt"
)
type Cat struct{
	Name string
	Age  int
	Color string
}
func main() {
	InitStruct()
}
func InitStruct(){
	var cat1 Cat
	fmt.Printf("init struct... %v\n",cat1)
}
func InitInstance(){
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 2
	cat1.Color = "白色"
	fmt.Printf("init instance cat1...%v\n",cat1)
}