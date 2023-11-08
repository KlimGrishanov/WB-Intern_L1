package main

import "fmt"

func someFunc() string {
	// т.к потом делаем слайс размером 100, не имеет смысла создавать 2^10, 2^7 покроет наши потребности
	v := createHugeString(1 << 7)
	// не используем глобальные переменные
	justString := v[:100]
	return justString
}

func createHugeString(size int) string {
	// создали
	runeArr := make([]rune, size)
	// вернули
	return string(runeArr)
}

func main() {
	str := someFunc()
	fmt.Println(str)
}
