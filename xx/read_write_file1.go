package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "hello.txt"
	outputFile := "world.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	fmt.Printf("##%s##\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)

	if err != nil {
		panic(err.Error())
	}
}
