package main

import "fmt"

// Работаем с bit операциями

func setBit(n int64, pos uint) int64 {
	return n | (int64(1) << pos)
}

func deleteBit(n int64, pos uint) int64 {
	return n &^ (int64(1) << pos)
}

func changeBitPos(n int64, pos uint, val bool) int64 {
	var result int64
	if val {
		result = setBit(n, pos)
	} else {
		result = deleteBit(n, pos)
	}
	return result
}

func main() {
	var i int64
	i = 10
	fmt.Println(changeBitPos(i, 2, true))
	fmt.Println(changeBitPos(i, 3, false))
}
