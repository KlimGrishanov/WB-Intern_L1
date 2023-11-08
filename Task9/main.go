package main

import (
	"fmt"
	"sync"
)

func FromArrayIntoChannel(in []int, out chan int, wg *sync.WaitGroup) {
	for i := 0; i < len(in); i++ {
		out <- in[i]
		wg.Done()
	}
}

func MultiplyChan(in chan int, out chan int, wg *sync.WaitGroup) {
	for el := range in {
		val := el * 2
		out <- val
		wg.Done()
	}
}

func StdoutChan(in chan int, wg *sync.WaitGroup) {
	for el := range in {
		fmt.Println(el)
		wg.Done()
	}
}

func main() {
	n := 12
	var arr []int
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	in := make(chan int)
	res := make(chan int)
	defer close(in)
	defer close(res)

	wg.Add(n * 3)
	go FromArrayIntoChannel(arr, in, &wg)
	go MultiplyChan(in, res, &wg)
	go StdoutChan(res, &wg)
	wg.Wait()
}
