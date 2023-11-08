package main

import (
	"context"
	"fmt"
	"time"
)

// Используем контекст для корректной остановки горутин

func ContextStopWorker(in chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received stop signal. Stopping goroutine.")
			close(in)
			return
		default:
			fmt.Println(<-in)
		}
	}
}

func Writer(out chan int) {
	i := 0
	for {
		out <- i
		i++
	}
}

func main() {
	in := make(chan int, 10)
	ctx, cancel := context.WithCancel(context.Background())
	defer close(in)

	for i := 0; i < 3; i++ {
		go ContextStopWorker(in, ctx)
	}
	go Writer(in)

	time.Sleep(5 * time.Second)
	cancel()
}
