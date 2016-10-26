package main

import (
	"regexp"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Println(" [ip_addr]")
		return
	}

	fmt.Println(flag.Arg(0), " is ip addr?\n", IsIP(flag.Arg(0)))
}

func IsIP(ip string)(b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
	//if m, _ := regexp.MatchString("^[[0-255]\\.]{3}[0-255]$", ip); !m {
		return false
	}
	return true
}