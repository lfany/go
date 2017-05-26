package main

import (
	"fmt"
	"github.com/lfany/go/trash/projs/c"
)

func main() {
	fmt.Println(rand.Random())
	rand.Seed(rand.Random())
	fmt.Println(rand.Random())

}
