package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {

	fName := "hello.zip"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Cannot Open %s, %v", os.Args[0], fName, err)
		os.Exit(1)
	}

	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		print("not zip\n")
		ff, _ := os.Open(fName)
		r = bufio.NewReader(ff)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("err: %v", err)
			os.Exit(0)
		}
		fmt.Print(line)
	}
}
