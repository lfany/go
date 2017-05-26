package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Echo(ws *websocket.Conn) {
	var err error

	for{
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Cannot receive!")
			break
		}

		fmt.Println("Received from client: #" + reply + "#")

		msg := "Received: #" + reply + "#"
		fmt.Println("Sending to client" + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cannot send!")
			break
		}
	}
}