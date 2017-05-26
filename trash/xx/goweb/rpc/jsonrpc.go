package main

import (
	"errors"
	"net/rpc"
	"net"
	"github.com/labstack/gommon/log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Divide(args *Args, quo *Quotient) error{
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func (t *Arith) Multiply(args *Args, reply *int)error  {
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkEEE(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkEEE(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		jsonrpc.ServeConn(conn)
	}
}
func checkEEE(i error) {
	if i != nil {
		log.Fatal(i)
	}
}
