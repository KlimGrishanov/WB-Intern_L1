package main

import "fmt"

type Human struct {
	height int
	weight int
}

func (h Human) getHeight() int {
	return h.height
}

func (h Human) getWeight() int {
	return h.weight
}

type Action struct {
	Human   Human
	SomeVar int
}

func (a Action) getHeight() int {
	return a.Human.getHeight()
}

func (a Action) getWeight() int {
	return a.Human.getWeight()
}

func main() {
	action := Action{Human: Human{183, 80}, SomeVar: 3}
	fmt.Println(action.getWeight())
	fmt.Println(action.getHeight())
}
