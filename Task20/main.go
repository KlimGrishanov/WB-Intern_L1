package main

import (
	"errors"
	"fmt"
)

// Решение основано на стеке

type Stack struct {
	len int
	arr []string
}

func (s *Stack) Push(value string) {
	s.arr = append(s.arr, value)
	s.len = len(s.arr)
}

func (s *Stack) Pop() (string, error) {
	if s.len <= 0 {
		return "", errors.New("Cant pop. No element in stack")
	}
	s.len -= 1
	return s.arr[s.len], nil
}

func createStack() Stack {
	return Stack{len: 0, arr: []string{}}
}

func reverseSentence(str string) (string, error) {
	stack := createStack()
	tempRuneArr := []rune(str)
	start, end, counter := 0, 0, 0

	for i := 0; i < len(tempRuneArr); i++ {
		if tempRuneArr[i] == ' ' || i == len(tempRuneArr)-1 {
			end = i
			if i == len(tempRuneArr)-1 {
				stack.Push(string(tempRuneArr[start : end+1]))
			} else {
				stack.Push(string(tempRuneArr[start:end]))
			}
			fmt.Println(stack.arr)
			start = i + 1
			counter++
		}
	}

	ans := ""
	for i := 0; i < counter; i++ {
		val, err := stack.Pop()
		if err != nil {
			return "", err
		}
		ans += val
		if i != counter-1 {
			ans += " "
		}
	}
	return ans, nil
}

func main() {
	str := "snow dog sun"
	reverseSen, err := reverseSentence(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(reverseSen)
}
