package main

import (
	"github.com/lfany/go/trash/gogo/structPack"
	"fmt"
)

func main() {
	struct1 := structPack.NewExpStruct()
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Println(struct1)
}
