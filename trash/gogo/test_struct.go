package main

import (
	"fmt"
	"github.com/lfany/go/gogo/structPack"
)

func main() {
	struct1 := structPack.NewExpStruct()
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Println(struct1)
}
