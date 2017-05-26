package main

import (
	"os"
	"fmt"
	"net/rpc"
	"github.com/labstack/gommon/log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		panic(err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Printf("Arith: %d / %d = %d -- %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
