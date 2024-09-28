package main

import "fmt"

func main() {
	a, b := 1, 2
	sum := 0
	for b < 4e6 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	fmt.Println(sum)
}
