package main

import (
	"fmt"
	"github.com/lfany/go/projs/stringutil"
	"runtime"
)

func main() {
	fmt.Printf("%s", stringutil.Reverse(runtime.Version()))
}
