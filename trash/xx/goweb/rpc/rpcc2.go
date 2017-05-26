package main

import (
	"github.com/labstack/gommon/log"
	"net/rpc"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func main() {
	service := "127.0.0.1:1234"

	client, err := rpc.Dial("tcp", service)
	checkEE(err)

	args := Args{16, 5}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkEE(err)
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	checkEE(err)

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
func checkEE(i error) {
	if i != nil {
		log.Fatal(i)
	}
}
