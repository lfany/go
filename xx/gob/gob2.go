package main

import (
	"os"
	"encoding/gob"
	"log"
)

type Address struct {
	Type 	string
	City	string
	Country	string
}

type VCard struct {
	FirstName	string
	LastName	string
	Address		[]*Address
	Remark 		string
}

var content string

func main() {
	pa := &Address{"private", "Acity", "JP"}
	wa := &Address{"work", "Bcity", "CN"}

	vc := VCard{"Jan", "ll", []*Address{pa, wa}, "none"}

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if(err != nil) {
		log.Println("Err int encode")
	}
}
