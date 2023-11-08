package main

import (
	"fmt"
	"sync"
)

// Остановка по окончанию задач от WaitGroup

func WaitGroupStopWorker(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println(<-in)
	}
}

func Writer(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		out <- i
		i++
	}
}

func main() {
	in := make(chan int, 10)
	var wg sync.WaitGroup
	defer close(in)

	wg.Add(2)

	for i := 0; i < 3; i++ {
		go WaitGroupStopWorker(in, &wg)
	}
	go Writer(in, &wg)

	wg.Wait()
}
