package main

import (
	"fmt"
)

func AsyncSquare(a int, out chan int) {
	out <- a * a
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	// Создаем канал и закрываем его по разрушению стека Main
	out := make(chan int)
	defer close(out)

	for i := 0; i < len(arr); i++ {
		// Запуск нескольких воркеров
		go AsyncSquare(arr[i], out)
	}

	// Считаем сумму из канала
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += <-out
	}
	fmt.Println(sum)
}
