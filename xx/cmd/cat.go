package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
	"io"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Args()[i])
		if err != nil {
			fmt.Printf("err :%v", err)
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			//fmt.Printf("err :%v", err)
			return
		}

		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}