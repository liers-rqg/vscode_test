package demo7

import "testing"

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 10 {
		t.Fatalf("执行错误，结果为%v\n", res)
	}
	t.Logf("执行正确")
}

func TestAddSelf(t *testing.T) {
	res := AddSelf(10)
	if res != 55 {
		t.Fatalf("执行结果为%v\n", res)
	}
	t.Logf("执行正确")
}
