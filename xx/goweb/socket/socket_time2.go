package main

import (
	"net"
	"fmt"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkE(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkE(err)
	for {
		conn, err := listener.Accept()
		//checkE(err)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	//defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
func checkE(err error) {
	if err != nil {
		panic(err)
	}
}
