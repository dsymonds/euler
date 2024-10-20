package main

import "fmt"

func main() {
	const (
		horiz = 20
		vert  = 20
	)

	// n choose k, where n=horiz+vert and k=horiz (or, equivalently, k=vert).
	// This is the multiplicative formula (https://en.wikipedia.org/wiki/Binomial_coefficient#Multiplicative_formula).
	const (
		n = horiz + vert
		k = horiz
	)
	// Compute carefully to keep this in integers, without blowing up too much.
	prod := 1
	var denom []int // factors of the denominator that haven't been used yet
	for i := 1; i <= k; i++ {
		prod *= (n + 1 - i)
		denom = append(denom, i)
		for j := len(denom) - 1; j >= 0; j-- { // see which denom factors can be used now
			if prod%denom[j] == 0 {
				prod /= denom[j]
				copy(denom[j:], denom[j+1:])
				denom = denom[:len(denom)-1]
			}
		}
	}
	if len(denom) > 0 {
		panic(fmt.Sprintf("Oh, dear. Didn't use up these denom factors: %v", denom))
	}
	fmt.Printf("horiz=%d vert=%d -> %d choose %d = %d\n", horiz, vert, n, k, prod)
}
