package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			n := i * j
			s := strconv.Itoa(n)
			ok := true
			for b := 0; b < len(s)/2; b++ {
				if s[b] != s[len(s)-b-1] {
					ok = false
					break
				}
			}
			if ok && n > max {
				// n is the new biggest palindrome
				max = n
			}
		}
	}
	fmt.Println("The largest palindrome made from the product of two 3-digit numbers is", max)
}
