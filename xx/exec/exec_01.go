package main

import (
	"os/exec"
	"fmt"
)

func main() {
	i, error := exec.Command("ls", "-al").Output();
	if error == nil {
		fmt.Println(string(i))
	} else {
		fmt.Println("exec error.", error)
	}
}
