package main

import (
	"flag"
	"fmt"
)

func PrimeFactorOf(n int) []int {
	ret := make([]int, 0)
	divider := 2
	for n > divider {
		for n%divider == 0 {
			n /= divider
			ret = append(ret, divider)
		}
		divider++
	}
	if n > 1 {
		ret = append(ret, n)
	}
	return ret
}

func main() {
	number := flag.Int("number", 0, "number")

	flag.Parse()
	fmt.Printf("Prime factor of %d, is %v\n", *number, PrimeFactorOf(*number))
}
