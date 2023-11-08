package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := new(big.Int), new(big.Int)
	a.SetString("4294954212", 10) // 2^32
	b.SetString("954729529", 10)  // 2^28

	tempVar := new(big.Int)
	fmt.Printf("Plus: %s\n", tempVar.Add(a, b).String())
	fmt.Printf("Minus: %s\n", tempVar.Sub(a, b).String())
	fmt.Printf("Multiply: %s\n", tempVar.Mul(a, b).String())
	fmt.Printf("Divide: %s\n", tempVar.Div(a, b).String())
}
