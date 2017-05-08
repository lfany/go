package main

import "fmt"

var sentinels = []string {
	"aa",
	"bb",
	"cc",
}

func main(){
	for _, i := range sentinels {
		fmt.Println(i)
	}
}