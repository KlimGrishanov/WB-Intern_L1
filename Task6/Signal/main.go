package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SignalStopWorker(in chan int, signal chan os.Signal) {
	for {
		select {
		case <-signal:
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
	sig := make(chan os.Signal, 1)
	defer close(in)
	defer close(sig)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < 3; i++ {
		go SignalStopWorker(in, sig)
	}
	go Writer(in)

	fmt.Println("Main goroutine is running. Press Ctrl+C to stop.")
	<-sig
	fmt.Println("Stopping main goroutine.")
}
