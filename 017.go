package main

import "fmt"

func main() {
	const limit = 1000
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += letterCount(i)
	}
	fmt.Println(sum)
}

// letterCount reports how many letters it would take to write out n
// in words (e.g. 115 as "one hundred and fifteen" is 20 letters).
func letterCount(n int) int {
	// We only need to handle 1-1000.
	if n < 1 || 1000 < n {
		panic("n out of range")
	}
	// Treat 1000 as a special case.
	if n == 1000 {
		return len("onethousand")
	}

	letters := 0
	// If n >= 100, we write "<n/100> hundred", possibly also adding "and <n%100>".
	if n >= 100 {
		letters += len(base[n/100]) + len("hundred")
		n %= 100
		if n == 0 {
			return letters
		}
		letters += len("and")
	}

	// We've reduced n to be in [1,99].

	// Handle awkward 10/11/12/teens.
	if 10 <= n && n <= 19 {
		letters += len(base[n])
		return letters
	}

	// Pull off tens place.
	if tens := n / 10 * 10; tens != 0 { // tens is 20/30/40/...
		letters += len(base[tens])
		n -= tens
	}

	// Finally, count the ones place.
	letters += len(base[n])

	return letters
}

var base = [...]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",

	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",

	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}
