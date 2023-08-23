package main

import "fmt"

func main() {
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//
	//name := rdb.Get(context.Background(), "name").Val()
	//fmt.Printf("name=%s\n", name)
	//fmt.Println(65536 >> 16)
	//fmt.Println(131072 >> 16)

	fmt.Println(1 ^ (1 << 16))
	fmt.Println(1 ^ (2 << 16))
}
