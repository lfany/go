package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type Address struct {
	Type 	string
	City 	string
	Country string
}

type VCard struct {
	FirstName 	string
	LastName 	string
	Address 	[]*Address
	Remark 		string
}

func main() {

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	js, _ := json.MarshalIndent(vc, "", "\t")
	fmt.Printf("Json Format: \n%s", js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	enc.SetEscapeHTML(true)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
