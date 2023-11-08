package main

import "fmt"

// Уберем возможность попадания нашего алгоритма в время выполнения O(n^2)
// Через выбор пивота из трех медиан, а не рандомом

func findMedianIndex(arr []int, start, end int) int {
	mid := start + (end-start)/2
	first, mid, last := arr[0], arr[mid], arr[end]
	if first < mid && first < last {
		if mid < last {
			return start + (end-start)/2
		} else {
			return end
		}
	} else if mid < first && mid < last {
		if first < last {
			return 0
		} else {
			return end
		}
	} else {
		if mid < first {
			return start + (end-start)/2
		} else {
			return 0
		}
	}
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := findMedianIndex(arr, low, high)
	arr[high], arr[pivot] = arr[pivot], arr[high]
	pivot = arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func main() {
	arr := [...]int{9, 2, 7, 10, 111, 46, 124, 51, 56, 62, 42, 33}
	fmt.Println(quickSort(arr[:], 0, len(arr)-1))
}
