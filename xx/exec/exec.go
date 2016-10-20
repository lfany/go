package main

import (
	"os"
	"fmt"
)

func main() {
	env := os.Environ()
	procAttr := os.ProcAttr{
		Env: env,
		Files:[]*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid, err := os.StartProcess("/usr/bin/ls", []string{"ls", "-l"}, &procAttr)
	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", *pid)
}
