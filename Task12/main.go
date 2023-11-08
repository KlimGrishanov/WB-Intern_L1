package main

import "fmt"

func main() {
	strArr := [...]string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for i := 0; i < len(strArr); i++ {
		set[strArr[i]] = struct{}{}
	}

	for v, _ := range set {
		fmt.Println(v)
	}
}
