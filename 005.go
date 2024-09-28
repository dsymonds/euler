package main

import "fmt"

func main() {
	// Find the smallest positive number that is evenly divisible by all of the numbers from 1 to N.
	const N = 20

	// By the fundamental theorem of arithmetic, we need prime powers that are the biggest that appear
	// in the factoring of the 1-N range.
	type prime struct {
		base  int // the prime itself
		power int // the exponent we need
	}
	var primes []*prime
	for i := 2; i <= N; i++ { // skip 1, since every integer is trivially divisible by it
		// Pull out all the prime factors.
		n := i
		for _, p := range primes {
			power := 0
			for n >= p.base {
				if n%p.base != 0 {
					break
				}
				n /= p.base
				power++
			}
			if power > p.power { // a new bigger power is required of this prime
				p.power = power
			}
		}
		if n != 1 {
			// We found a new prime.
			primes = append(primes, &prime{base: n, power: 1})
		}
	}

	// Compute the product.
	x := 1
	for _, p := range primes {
		for i := 0; i < p.power; i++ {
			x *= p.base
		}
	}
	fmt.Printf("Smallest positive number that is evenly divisible by all of the numbers from 1-%d is %d\n", N, x)
}
