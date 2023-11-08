package main

import "fmt"

func reverseStr(in string) string {
	tempRuneArr := []rune(in)
	resRuneArr := make([]rune, len(tempRuneArr))
	for i := len(tempRuneArr) - 1; i >= 0; i-- {
		resRuneArr[len(tempRuneArr)-1-i] = tempRuneArr[i]
	}
	return string(resRuneArr)
}

func main() {
	str := "главрыба"
	fmt.Println(reverseStr(str))
}
