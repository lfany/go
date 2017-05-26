package main

import (
	//"time"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}
func getData(strings chan string) {
	var input string

	for {
		input = <-strings
		fmt.Printf("%s ", input)
	}
}
func sendData(strings chan string) {
	strings <- "hello1"
	strings <- "hello2"
	strings <- "hello3"
	strings <- "hello4"
	strings <- "hello5"
	strings <- "hello6"
}
