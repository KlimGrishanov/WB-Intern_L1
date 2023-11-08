package main

import "fmt"

func binSearch(arr []int, start, end, val int) int {
	for end-start > 0 {
		mid := start + (end-start)/2
		if val == arr[mid] {
			return mid
		} else if val > arr[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(binSearch(arr[:], 0, len(arr)-1, 10))
	fmt.Println(binSearch(arr2[:], 0, len(arr2)-1, 11))
}
