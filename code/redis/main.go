package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func open() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func main() {
	fmt.Println("Go Regis")

	rdb := open()
	defer rdb.Close()

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = rdb.Set("name", "risa", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	var name string
	name, err = rdb.Get("name").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}