package main

import (
	"gopkg.in/redis.v5"
	"fmt"
	"log"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"", // no password set
		DB: 		0, //use default DB
	})

	pong, err := client.Ping().Result()
	checkE(err)

	fmt.Println("pong: ", pong)

	err = client.Set("key", "value", 0).Err()
	checkE(err)

	val, err := client.Get("key").Result()

	fmt.Println("val: ", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 doesnot exists")
	} else if err != nil {
		panic(err)
	}else {
		fmt.Println("key2: ", val2)
	}


}

func checkE(err error) {
	if err != nil {
		log.Fatal(err)
	}
}