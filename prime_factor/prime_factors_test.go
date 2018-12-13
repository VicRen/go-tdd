package prime_factor

import (
	"reflect"
	"testing"
)

func TestPrimeFactorOf(t *testing.T) {
	tt := []struct {
		name string
		in   int
		out  []int
	}{
		{"Prime factor of one", 1, []int{}},
		{"Prime factor of two", 2, []int{2}},
		{"Prime factor of three", 3, []int{3}},
		{"Prime factor of four", 4, []int{2, 2}},
		{"Prime factor of five", 5, []int{5}},
		{"Prime factor of six", 6, []int{2, 3}},
		{"Prime factor of seven", 7, []int{7}},
		{"Prime factor of eight", 8, []int{2, 2, 2}},
		{"Prime factor of nine", 9, []int{3, 3}},
		{"Prime factor of ten", 10, []int{2, 5}},
		{"Prime factor of a big number", 2 * 2 * 3 * 5 * 7 * 11 * 13 * 17 * 23,
			[]int{2, 2, 3, 5, 7, 11, 13, 17, 23}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := PrimeFactorOf(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("Prime factor of one should be %v; got %v", tc.out, out)
			}
		})
	}
}
