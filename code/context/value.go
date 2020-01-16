package main

import (
	"fmt"
	"context"
	"time"
)

type favContextKey string

func getName(ctx context.Context, k favContextKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("Found name:", v)
		return
	}
	fmt.Println("Name is not found")
}

func main() {
	fmt.Println("context package 2")

	k := favContextKey("name")

	cancelCxt := context.WithValue(context.Background(), k, "[cancelContext]")
	timeoutCxt := context.WithValue(context.Background(), k, "[timeoutContext]")
	deadlineCxt := context.WithValue(context.Background(), k, "[deadlineContext]")

	go getName(cancelCxt, k)
	go getName(timeoutCxt, k)
	go getName(deadlineCxt, k)

	time.Sleep(time.Second)
}