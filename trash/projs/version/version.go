package main

import (
	"fmt"
	"github.com/lfany/go/trash/projs/stringutil"
	"runtime"
)

func main() {
	fmt.Printf("%s", stringutil.Reverse(runtime.Version()))
}
