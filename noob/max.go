package main

import "fmt"

var a, b = 100, 200

func main() {

	fmt.Println("Max: ", max(a, b))
}

func max(a, b int) int {
	if a > b {
		b = a
	}
	return b
}
