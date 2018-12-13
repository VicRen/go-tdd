package prime_factor

func PrimeFactorOf(n int) []int {
	var ret = make([]int, 0)
	divisor := 2
	for divisor < n {
		for n%divisor == 0 {
			n /= divisor
			ret = append(ret, divisor)
		}
		divisor ++
	}
	if n > 1 {
		ret = append(ret, n)
	}
	return ret
}
