package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("error!!")
		return
	}

	fmt.Printf("%s", input)

	switch input {
	case "abc\n":
		fallthrough
	case "bcd\n":
		fallthrough
	case "zxc\n":
		fmt.Println("Shadow Field")
	case "lol\n":
		fmt.Println("Dota")
	default:
		fmt.Println("Good by!")
	}

}
