package main

import "fmt"

func deleteElem(arr []int, indexPos int) []int {
	newArr := append(arr[0:indexPos], arr[indexPos+1:]...)
	return newArr
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	pos := 2
	newArr := deleteElem(arr[:], pos)
	fmt.Println(newArr)
}
