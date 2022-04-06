package demo5

import "strings"

// 闭包
//累加器
func My_closed_func() func(x int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func MakeSufix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
