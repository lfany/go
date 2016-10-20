package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting...")
	listener, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Println("err", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err accept", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err read", err.Error())
			return
		}
		fmt.Printf("Read: %v", string(buf))
	}
}
