package demo7

//单元测试
func AddUpper(num int) int {
	res := 0
	for i := 0; i < num; i++ {
		res++
	}
	return res
}

func AddSelf(num int) int {
	res := 0
	for i := 0; i < num; i++ {
		res += i
	}
	return res
}
