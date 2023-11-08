package main

import "fmt"

func main() {
	arr1 := []int{2, 4, 3, 5, 6, 7}
	arr2 := []int{9, 2, 7, 6}
	var intersectionArr []int

	firstSet := make(map[int]struct{})
	secondSet := make(map[int]struct{})
	intersection := make(map[int]struct{})

	// Готовим данные (переводим в мапку)
	for i := 0; i < len(arr1); i++ {
		firstSet[arr1[i]] = struct{}{}
	}
	for i := 0; i < len(arr2); i++ {
		secondSet[arr2[i]] = struct{}{}
	}

	// Заполняем наше множество пересечений из первого мн-ва
	for i := 0; i < len(arr1); i++ {
		_, isExistInFirst := firstSet[arr1[i]]
		_, isExistInSecond := secondSet[arr1[i]]
		if isExistInSecond && isExistInFirst {
			intersection[arr1[i]] = struct{}{}
		}
	}

	// Заполняем наше множество пересечений из второго мн-ва
	for i := 0; i < len(arr2); i++ {
		_, isExistInFirst := firstSet[arr2[i]]
		_, isExistInSecond := secondSet[arr2[i]]
		if isExistInSecond && isExistInFirst {
			intersection[arr2[i]] = struct{}{}
		}
	}

	// Преобразуем множество пересечений в массив
	for v, _ := range intersection {
		intersectionArr = append(intersectionArr, v)
	}

	// Выводим массив пересечений
	fmt.Println(intersectionArr)
}
