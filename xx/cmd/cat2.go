package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat2(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))

		if err != nil || f == nil {
			fmt.Printf("err: %v", err)
			os.Exit(1)
		}
		cat2(f)
		f.Close()
	}
}

func cat2(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte

	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "err: %s\n", err)
			os.Exit(1)
		case nr == 0:
			//	EOF
			return
		case nr > 0:
			if nw, err := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "err :%v", err)
			}

		}
	}
}

func write() {
	fmt.Fprintf(os.Stdout, "hello")
}
