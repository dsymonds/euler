package main

import "fmt"

func main() {
	// Initial set of primes:
	primes := []int{2, 3, 5, 7}
	const limit int = 2e6

	// Start with the sum of the initial set.
	sum := 0
	for _, p := range primes {
		sum += p
	}
	// Keep adding new primes to the list.
outer:
	for j := primes[len(primes)-1] + 1; j < limit; j++ {
		// See if j is prime.
		for _, p := range primes {
			if j%p == 0 {
				continue outer
			}
		}
		// j does not have any smaller prime as a factor, so it is prime.
		primes = append(primes, j)
		sum += j
	}

	fmt.Printf("Sum of primes under %d is %d\n", limit, sum)
}
