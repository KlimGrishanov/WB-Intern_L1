package main

import (
	"context"
	"fmt"
	"time"
)

func writer(out chan int, ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out <- i
			i++
		}
	}
}

func worker(in chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-in)
		}
	}
}

func main() {
	channel := make(chan int, 10)
	ctx := context.Background()
	defer close(channel)

	go worker(channel, ctx)
	go writer(channel, ctx)

	select {
	case <-time.After(1 * time.Second):
		ctx.Done()
		fmt.Println("Time Out!")
	}
}
