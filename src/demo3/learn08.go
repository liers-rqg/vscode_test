package demo3

//定义一个单链表
type IntLinkList struct {
	Value int
	Tail  *IntLinkList
}

func (list *IntLinkList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

//定义一个已有数据类型
type Integer int

func (integer *Integer) Increace() int {
	(*integer)++
	return int(*integer)
}
