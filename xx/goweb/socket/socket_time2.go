package main

import (
	"net"
	"fmt"
	"time"
	"strings"
	"strconv"
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
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(string(request))

		if readLen == 0 {
			fmt.Println("##:" + string(readLen))
			break // connection already closed by client
		}else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		}else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
	}
	request = make([]byte, 128) // clear last read content
	//daytime := time.Now().String()
	//conn.Write([]byte(daytime))
}
func checkE(err error) {
	if err != nil {
		panic(err)
	}
}
