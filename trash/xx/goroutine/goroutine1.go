package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In Main")
	go longWait()
	go shortWait()

	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("End of main")
}
func shortWait() {
	fmt.Println("Begin short")
	time.Sleep(2 * 1e9)
	fmt.Println("end of short")
}
func longWait() {
	fmt.Println("Begin long")
	time.Sleep(5 * 1e9)
	fmt.Println("end of long")
}
