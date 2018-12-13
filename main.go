package main

import (
	"flag"
	"fmt"
	"github.com/VicRen/go-tdd/prime_factor"
)

func main() {
	number := flag.Int("number", 0, "number")

	flag.Parse()
	fmt.Printf("Prime factor of %d, is %v", *number, prime_factor.PrimeFactorOf(*number))
}
