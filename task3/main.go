package main

import (
	"fmt"
)

func AsyncSquare(a int, out chan int) {
	out <- a * a
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	out := make(chan int)
	defer close(out)

	for i := 0; i < len(arr); i++ {
		go AsyncSquare(arr[i], out)
	}

	var sum int
	for i := 0; i < len(arr); i++ {
		sum += <-out
	}
	fmt.Println(sum)
}
