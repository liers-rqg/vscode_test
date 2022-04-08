package demo9

import (
	"fmt"
	"sync"
)

var intchan chan int

//go本身提供了一个waitgroup用于判断协程是否执行完毕
//开启一个协程add
//协程执行完毕后done
//所有的add被done后才会执行wait
var wg sync.WaitGroup

func WriteRoutine(n int) {
	intchan <- n
	wg.Done()
}
func ReadRoutine() {
	fmt.Println(<-intchan)
	wg.Done()
}

//func fakemain() {
// intchan = make(chan int, 50)
// for i := 0; i < 50; i++ {
// 	wg.Add(1)
// 	go WriteRoutine(i)
// }
// for i := 0; i < 50; i++ {
// 	wg.Add(1)
// 	go ReadRoutine()
// }
// wg.Wait()

// close(intchan)
// primeChan := make(chan int, 100)
// go putRoutine(500, primeChan)
// go isPrimeRoutine(primeChan)

//}
