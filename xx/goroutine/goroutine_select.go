package main

import (
	"time"
	"fmt"
)


var (
	i1 int
	i2 int
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)

	go suck12(ch1, ch2)

	time.Sleep(1e9)
	fmt.Printf("counts: i1-%d\ti2-%d\n", i1, i2)
}
func pump2(ints chan int) {
	for i := 0;  ; i++ {
		ints <- i * 2
	}
}
func pump1(ints chan int) {
	for i := 0;  ; i++ {
		ints <- i + 5
	}
}

func suck12(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v:= <-ch1:
			i1++
			fmt.Printf("ch1: %d\n", v)
		case v := <-ch2:
			i2++
			fmt.Printf("ch2: %d\n", v)
		}
	}
}

