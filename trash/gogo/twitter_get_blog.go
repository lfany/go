package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	u := "https://blog.twitter.com/2017/the-infrastructure-behind-twitter-scale"
	r, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	log.Print(r)
	io.Copy(os.Stdout, r.Body)
}
