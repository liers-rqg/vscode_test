package demo9

import (
	"fmt"
	"sync"
)

func isprime(n int) bool {
	if n <= 2 {
		return false
	}
	res := true
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return res
}

func PutRoutine(n int, primeChan chan int, wg *sync.WaitGroup) {
	for i := 1; i < n; i++ {
		primeChan <- i
	}
	defer wg.Done()
	close(primeChan)

}
func IsPrimeRoutine(primeChan chan int, numChan chan int, wg *sync.WaitGroup) {
	for {
		k, v := <-primeChan
		if !v {
			fmt.Println(v)
			break
		}
		if isprime(k) {
			numChan <- k
		}
	}
	defer wg.Done()
}

func GetPrime(primeChan chan int, wg *sync.WaitGroup) {

	for {
		k, v := <-primeChan
		if !v {
			break
		}
		fmt.Println(k)
	}
	defer wg.Done()
}
