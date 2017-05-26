package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

var (
	content string
	vc      VCard
)

func main() {
	file, _ := os.Open("vcard.gob")
	defer file.Close()

	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Err in gob")
	}
	fmt.Println(vc)
	ll := vc.Address

	for i := 0; i < len(ll); i++ {
		fmt.Println(*ll[i])
	}
	fmt.Println(vc.Address)

}
