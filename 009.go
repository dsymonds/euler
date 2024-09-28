package main

import (
	"fmt"
	"math"
)

func main() {
	const target = 1000
aLoop:
	for a := 1; ; a++ {
		// b < c, and c = 1000-(a+b),
		// so b < 1000-(a+b),
		// or 2b < 1000-a
	bLoop:
		for b := a + 1; b < (1000-a)/2; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2))) // rounded c
			switch sum := a + b + c; {
			case sum > target:
				// Gone too high. Next a.
				continue aLoop
			case sum < target:
				// Too low. Next b.
				continue bLoop
			}
			// Check if c is integral.
			if c*c != c2 {
				// Not a Pythagorean triplet.
				continue
			}
			fmt.Printf("a=%d b=%d c=%d; a+b+c=%d; abc=%d\n", a, b, c, a+b+c, a*b*c)
			return
		}
	}
}
