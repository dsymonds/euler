package main

import (
	"fmt"
	"math/big"
)

func main() {
	// What's the sum of the (decimal) digits of 2^1000?
	const power = 1000

	var n big.Int
	n.SetBit(&n, power, 1)
	s := n.Text(10) // render as a decimal string

	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	fmt.Printf("Digit sum of 2^%d is %d\n", power, sum)
}
