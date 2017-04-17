package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.71.178:6379",
		Password: "00",
		DB:       2,
	})

	pong, err := client.Set("aaa", "你好", 0).Result()
	fmt.Println(pong, err)

	pong, err = client.Get("aaa").Result()
	if err == redis.Nil {
		fmt.Println("##" + err.Error() + "##")
	}else{
		fmt.Println(pong, err)
	}

}
