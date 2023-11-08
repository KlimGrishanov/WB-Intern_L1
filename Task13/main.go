package main

import "fmt"

func swap(x, y *int) {
	if x == y {
		return
	}
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

func main() {
	x, y := 300, 400
	swap(&x, &y)
	fmt.Println(x, y)
}
