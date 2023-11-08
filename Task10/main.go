package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for i := 0; i < len(data); i++ {
		key := int(data[i] / 10)
		groups[key*10] = append(groups[key*10], data[i])
	}

	for v, k := range groups {
		fmt.Printf("Key: %d; Value: %f\n", v, k)
	}
}
