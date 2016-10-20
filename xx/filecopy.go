package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	i, err := CopyFile("mm.txt", "hello.txt")
	if err != nil {
		fmt.Printf("err : %v", err)
		return
	}
	fmt.Println("Copy done!", i)
}
func CopyFile(dstFile string, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("err : %v", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_TRUNC, 0)
	if err != nil {
		fmt.Printf("err : %v", err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}


