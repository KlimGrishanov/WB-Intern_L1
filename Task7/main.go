package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(syncMap *sync.Map) {
	i := 0
	for {
		syncMap.Store(i, i*i)
		time.Sleep(time.Millisecond * 100)
		i++
	}
}

func main() {
	var syncMap sync.Map

	go Worker(&syncMap)
	time.Sleep(time.Second)

	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})
}
