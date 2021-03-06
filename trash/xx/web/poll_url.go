package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com/",
	"http://163.com/",
	"http://blog.golang.org/",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Print(url, ": ", resp.Status)
	}
}
