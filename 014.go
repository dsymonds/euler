package main

import "fmt"

var memo []int // 0 = unchecked

func Collatz(n int) int {
	orig := n
	i := 1
	for n > 1 {
		if len(memo) > n && memo[n] > 0 {
			i += memo[n] - 1
			break
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		i++
	}
	for len(memo) <= orig {
		memo = append(memo, 0)
	}
	memo[orig] = i
	return i
}

func main() {
	max, maxChain := 0, 0
	for i := 1; i <= 1e6; i++ {
		chain := Collatz(i)
		if chain > maxChain {
			max, maxChain = i, chain
		}
	}
	fmt.Println(max, "has a chain of length", maxChain)
}
