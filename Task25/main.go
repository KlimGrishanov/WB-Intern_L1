package main

import (
	"fmt"
	"time"
)

func CustomSleep(sec int64) {
	now := time.Now().Unix()
	fmt.Println(now)
	for time.Now().Unix()-now != sec {
		// Just Wait
	}
	fmt.Println("After")
}

func main() {
	fmt.Println("Start")
	CustomSleep(5)
	fmt.Println("After 5 sec")
}
