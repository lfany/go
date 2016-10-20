package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	var str string

	if len(os.Args) > 1 {
		str += strings.Join(os.Args[1:], "\n")
	}

	fmt.Println(str)
}
