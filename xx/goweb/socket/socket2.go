package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	var service string

	if len(os.Args) != 2 {
		//fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		//os.Exit(1)
		service = "127.0.0.1:1200"
	} else {
		service = os.Args[1]
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	_, err = conn.Write([]byte("timestamp"))
	checkErr(err)
	//result, err := ioutil.ReadAll(conn)
	//checkErr(err)
	//fmt.Println(string(result))

	readServer(conn)
	os.Exit(0)

}
func readServer(conn *net.TCPConn) {
	defer conn.Close()
	result := make([]byte, 128)
	readLen, err := conn.Read(result)
	checkErr(err)
	fmt.Println(string(result), "##", readLen)

	if readLen == 0 {
		fmt.Println("Server Closed")
	} else {
		fmt.Println("xxx:", string(result[:readLen]), "xxx")
	}
	result = make([]byte, 128)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
