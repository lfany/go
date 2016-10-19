package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	//ms := new(struct1)   ms => *struct1
	var ms struct1 // ms => struct1
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Hello"

	fmt.Printf("Int: %d\n", ms.i1)
	fmt.Printf("Float32: %f\n", ms.f1)
	fmt.Printf("String: %s\n", ms.str)

	fmt.Println(ms)
	fmt.Printf("%v\n", ms)
}
