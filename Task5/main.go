package main

import (
	"fmt"
	"time"
)

func writer(out chan int) {
	i := 0
	for {
		out <- i
		i++
	}
}

func worker(in chan int) {
	for {
		fmt.Println(<-in)
	}
}

func main() {
	channel := make(chan int, 10)
	defer close(channel)

	go worker(channel)
	go writer(channel)

	select {
	case <-time.After(1 * time.Second):
		close(channel)
		fmt.Println("Time Out!")
	}
}
