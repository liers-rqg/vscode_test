package main

import (
	"fmt"
	"sync"
	"time"
)

var myMap = make(map[int]int)
var lock sync.Mutex

func cal(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}
func calRoutine(n int) {
	temp := cal(n)
	lock.Lock()
	myMap[n] = temp
	lock.Unlock()
}
func routine() {
	for i := 0; i < 10; i++ {
		//fmt.Println("hello routine: " + strconv.Itoa(i))
		//time.Sleep(time.Second)
	}

}

func main() {
	go routine() //开启一个协程
	for i := 0; i < 10; i++ {
		//fmt.Println("hello golang: " + strconv.Itoa(i))
		//time.Sleep(time.Second)
	}
	for i := 1; i <= 10; i++ {
		go calRoutine(i)
	}

	time.Sleep(1 * time.Second)
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("%d : %d \n", k, v)
	}
	lock.Unlock()
	channelExample()
}
