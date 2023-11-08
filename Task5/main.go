package main

import (
	"context"
	"fmt"
	"time"
)

// Используем контекст для корректной остановки горутин

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

	// Создаем воркеров
	go worker(channel, ctx)
	go writer(channel, ctx)

	// Создаем задержку
	<-time.After(1 * time.Second)
	// Ждем конца
	ctx.Done()
	fmt.Println("Time Out!")
}
