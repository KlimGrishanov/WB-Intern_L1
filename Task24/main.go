package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func createPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetDistance(toPoint *Point) float64 {
	dist := math.Sqrt(math.Pow(math.Abs(float64(toPoint.x-p.x)), 2) + math.Pow(math.Abs(float64(toPoint.y-p.y)), 2))
	return dist
}

func main() {
	Point1 := createPoint(1, 2)
	Point2 := createPoint(4, 6)
	fmt.Println(Point2.GetDistance(Point1))
}
