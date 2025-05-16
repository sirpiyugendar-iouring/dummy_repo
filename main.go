package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ch := rdb.Subscribe("new").Channel()
	for {
		msg := <-ch
		fmt.Println(msg.Payload)
	}
}
