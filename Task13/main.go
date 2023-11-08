package main

import "fmt"

func main() {
	x, y := 300, 400
	x, y = y, x
	fmt.Println(x, y)
}
