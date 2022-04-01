package demo2

import(
	"fmt"
)
var Num = TypeFunc
func main02(){
	num := Calculator(1.2,1.3,'-')
	fmt.Println(num)
}

func Calculator(n1 float64,n2 float64,opt byte) float64 {
	var res float64
	switch opt {
	case '+': res = n1 + n2
	case '-': res = n1 - n2
	case '*': res = n1 * n2
	case '/': res = n1 / n2
	default: {
		fmt.Println("输入错误")
		res = 0.0
	}
	}
	return res
}
func MultyReturn()(num int,num2 int){
	return 0,1
}
func TypeFunc(n1 int, n2 int)  int{
	return n1+n2
}