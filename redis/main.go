package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	name := rdb.Get(context.Background(), "name").Val()
	log.Println(name)
}
