package main

import (
	"fmt"
	"github.com/lfany/go/projs/c"
)

func main() {
	fmt.Println(rand.Random())
	rand.Seed(10)
	fmt.Println(rand.Random())

}
