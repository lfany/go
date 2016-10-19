package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)
func main() {
	inputFile, inputErr := os.Open("hello.txt")
	if inputErr != nil {
		fmt.Printf("Error: %s\n", inputErr)
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for{
		inputString, err := inputReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Printf("##%s##\n", inputString)
	}
}
