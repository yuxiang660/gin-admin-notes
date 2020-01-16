package main

import (
	"fmt"
	"context"
	"time"
)

func loop(ctx context.Context, name string) {
Loop:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println(name, "is running...")
		case <-ctx.Done():
			fmt.Println(name, "is done.")
			break Loop
		}
	}
}

func main() {
	fmt.Println("context package")

	cancelCxt, cancel := context.WithCancel(context.Background())
	timeoutCxt, timeoutCancel := context.WithTimeout(context.Background(), 3 * time.Second)
	deadlineCxt, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(4 * time.Second))

	defer func() {
		cancel()
		timeoutCancel()
		deadlineCancel()
	}()

	go loop(cancelCxt, "[cancelContext]")
	go loop(timeoutCxt, "[timeoutContext]")
	go loop(deadlineCxt, "[deadlineContext]")

	time.Sleep(time.Second)
	cancel()
	<- cancelCxt.Done()
	fmt.Println("After one second, cancel the cancel context actively.")

	<- timeoutCxt.Done()
	fmt.Println("After three seconds, the timeout context is done.")

	<- deadlineCxt.Done()
	fmt.Println("After four seconds, the timeout context is done.")
}