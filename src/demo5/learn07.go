package demo5

import "fmt"

// defer
func Func_defer(n ...int) {
	len := len(n)
	for i := 0; i < len; i++ {
		defer fmt.Println(n[i])
	}
	println("done...")
}
