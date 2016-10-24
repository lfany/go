package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.71.66:66")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.168.71.66:66")
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "192.168.71.66:66")
	checkConnection(conn, err)
}
func checkConnection(conn net.Conn, i error) {
	if i != nil {
		fmt.Printf("error %v connecting", i)
		os.Exit(1)
	}
	fmt.Printf("conn : %v\n", conn)
}
