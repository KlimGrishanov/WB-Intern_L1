package main

import (
	"fmt"
	"strings"
)

func checkUniq(str string) bool {
	uniqMap := make(map[rune]struct{})
	str = strings.ToLower(str)
	tempRuneArr := []rune(str)

	// Создаем Set на основе мапы
	for i := 0; i < len(tempRuneArr); i++ {
		if _, isExist := uniqMap[tempRuneArr[i]]; isExist {
			return false
		}
		uniqMap[tempRuneArr[i]] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(checkUniq("abCdefAaf"))
	fmt.Println(checkUniq("aabcd"))
	fmt.Println(checkUniq("abcd"))
}
