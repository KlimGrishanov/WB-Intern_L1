package main

import (
	"fmt"
	"sync"
)

func AsyncSquare(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a * a)
}

func main() {
	// Про порядок ничего не сказано в тз, поэтому не сохраняю его
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go AsyncSquare(arr[i], &wg)
	}

	wg.Wait()
}
