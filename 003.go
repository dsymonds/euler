package main

import "fmt"

func main() {
	const N = 600851475143

	n := N   // reduced by factoring out primes
	max := 0 // largest prime factor, or zero if we haven't seen one
	consider := func(p int) {
		// pull out as many multiples of p as we can
		for n%p == 0 {
			max = p
			n /= p
		}
	}

	// Initial set of primes:
	primes := []int{2, 3, 5, 7, 11, 13}
	for _, p := range primes {
		consider(p)
		if n < p {
			break
		}
	}

	// Keep adding new primes to the list, and removing each one from n as we go.
seedLoop:
	for j := primes[len(primes)-1] + 1; j*j < N; j++ {
		// See if j is prime.
		for _, p := range primes {
			if j%p == 0 {
				continue seedLoop
			}
		}
		// j does not have any smaller prime as a factor, so it is prime.
		primes = append(primes, j)
		consider(j)
		if n < j {
			break
		}
	}

	fmt.Println("The largest prime factor of", N, "is", max)
}
