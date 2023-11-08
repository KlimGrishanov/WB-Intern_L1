package main

import "fmt"

func WhatType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("Type of %d is int\n", a)
	case string:
		fmt.Printf("Type of '%s' is string\n", a)
	case bool:
		fmt.Println("Type is bool")
	case chan int:
		fmt.Println("Type is int chan")
	case chan string:
		fmt.Println("Type is string chan")
	case chan bool:
		fmt.Println("Type is bool chan")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var test interface{}
	test = true
	WhatType(test)
}
