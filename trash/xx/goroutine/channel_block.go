package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	<-ch1
	time.Sleep(115 * 1e9)

}
func pump(ints chan int) {
	for i := 0; ; i++ {
		ints <- i
	}
}

func suck(ints chan int) {
	for {
		fmt.Println(<-ints)
		//time.Sleep(1 * 1e9)
	}
}
