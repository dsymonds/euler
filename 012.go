package main

import "fmt"

func main() {
	// Go through the sequence of triangular numbers, counting the divisors of each one.
	const ndivisors = 500
	for i, n := 1, 0; ; i++ { // n is the triangular number
		n += i

		// Count the factors; 1 and n are trivially included. (this is wrong for n=1)
		divs := 2
		for x := 2; ; x++ {
			if x*x > n {
				break
			}
			if n%x != 0 {
				continue
			}
			// If x*x = n, only count it once; otherwise, it counts for two (x and n/x).
			if x*x == n {
				divs++
			} else {
				divs += 2
			}
		}
		if divs > ndivisors {
			fmt.Printf("%d has over %d divisors\n", n, ndivisors)
			return
		}
	}
}
