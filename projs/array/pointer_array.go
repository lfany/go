package main

import "fmt"

func f(a [3]int) {
	a[0] = 100
	fmt.Println(a)
}
func fp(a *[3]int) {
	a[0] = 100
	fmt.Println(a)
}

func main() {
	var ar [3]int
	fmt.Println(ar)
	f(ar)
	fmt.Println(ar)
	fp(&ar)
	fmt.Println(ar)
}
