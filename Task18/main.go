package main

import (
	"fmt"
	"sync"
)

// Обвешиваем мьютексами в вэйтгруппом чтобы ждать выполнения тасок и не ловить гонки

type Counter struct {
	wg      *sync.WaitGroup
	mut     *sync.Mutex
	counter int
}

func (c *Counter) getCounter() int {
	return c.counter
}

// Защищаемся от гонок
func (c *Counter) Increment() {
	defer c.mut.Unlock()
	defer c.wg.Done()
	c.mut.Lock()
	c.counter++
}

// Конструктор
func createCounter() Counter {
	return Counter{mut: &sync.Mutex{}, wg: &sync.WaitGroup{}, counter: 0}
}

func main() {
	c := createCounter()
	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		go c.Increment()
	}
	c.wg.Wait()
	fmt.Println(c.getCounter())
}
