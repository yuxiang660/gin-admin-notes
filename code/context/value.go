package main

import (
	"fmt"
	"context"
	"time"
)

type (
	cancelContextKey struct{}
	timeoutContextKey struct{}
	deadlineContextKey struct{}
)

func getName(ctx context.Context, k interface{}) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("Found name:", v)
		return
	}
	fmt.Println("Name is not found")
}

func main() {
	fmt.Println("context package 2")


	cancelCxt := context.WithValue(context.Background(), cancelContextKey{}, "[cancelContext]")
	timeoutCxt := context.WithValue(context.Background(), timeoutContextKey{}, "[timeoutContext]")
	deadlineCxt := context.WithValue(context.Background(), deadlineContextKey{}, "[deadlineContext]")

	go getName(cancelCxt, cancelContextKey{})
	go getName(timeoutCxt, timeoutContextKey{})
	go getName(deadlineCxt, deadlineContextKey{})

	time.Sleep(time.Second)
}