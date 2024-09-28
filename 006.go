package main

import "fmt"

func main() {
	const N = 100

	sumOfSquares := 0
	sum := 0
	for i := 1; i <= N; i++ {
		sumOfSquares += i * i
		sum += i
	}
	squareOfSum := sum * sum

	fmt.Println("squareOfSum - sumOfSquares =", squareOfSum-sumOfSquares)
}
