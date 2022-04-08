package main

import "fmt"

func channelExample() {
	intChannel := make(chan int, 10)
	intChannel <- 10
	intChannel <- 11
	intChannel <- 12
	for i := 0; i < len(intChannel); i++ {
		fmt.Println(<-intChannel)
	}
}
