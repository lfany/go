package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Thisis an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" has prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
