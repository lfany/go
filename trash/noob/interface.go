package main

import "fmt"

type interface_name interface {
	interface_methord()
}

type struct_name struct {
	name string
}
type struct2_name struct{}

func (xxx *struct_name) call() {
	fmt.Println("Struct Name: ", xxx.name)
}

func main() {
	var xxx struct_name

	xxx = struct_name{
		name: "xxxxx",
	}

	xxx.call()
}
