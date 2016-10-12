package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	for z := 0.0000001; z - (z- (z*z -x)/2z) <= 0.0000001; z += 0.0000001 {
			return z;
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
