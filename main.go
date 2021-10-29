package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopherCat/contents"
)

const RootDir = "/Volumes/Seagate Expansion Drive/comics/on-deck"

func main() {

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})


	val, err := rdb.Do(ctx, "ping").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println("Redis: ", val)


	msg := contents.Walk(RootDir)
	fmt.Println(msg)
}
