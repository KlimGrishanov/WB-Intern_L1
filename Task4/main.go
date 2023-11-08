package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func worker(in chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(in)
			return
		default:
			fmt.Println(<-in)
		}
	}
}

func main() {
	var numOfWorker int
	_, _ = fmt.Scanf("%d", &numOfWorker)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	in := make(chan int, 10)
	defer close(in)

	for i := 0; i < numOfWorker; i++ {
		go worker(in, ctx)
	}

	i := 0
	for {
		in <- i
		i++
	}
}
