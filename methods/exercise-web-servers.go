package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	A string
	B string
	C string
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	// var h Hello
	http.Handle("/hello", String("hello"))
	http.Handle("/memeda", &Struct{"aa", ":", "xxxxxx"})
	err := http.ListenAndServe("192.168.71.228:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
