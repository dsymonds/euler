package main

import "fmt"

func main() {
	// Initial set of primes:
	primes := []int{2, 3, 5, 7, 11, 13}
	const nth = 10001

	// Keep adding new primes to the list until we get the nth.
outer:
	for j := primes[len(primes)-1] + 1; len(primes) < nth; j++ {
		// See if j is prime.
		for _, p := range primes {
			if j%p == 0 {
				continue outer
			}
		}
		// j does not have any smaller prime as a factor, so it is prime.
		primes = append(primes, j)
	}

	fmt.Printf("The %dth prime is %d\n", nth, primes[nth-1])
}
