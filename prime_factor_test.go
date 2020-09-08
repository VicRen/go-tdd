package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactorOf(t *testing.T) {
	tt := []struct {
		name   string
		input  int
		output []int
	}{
		{
			"prime factor of 1",
			1,
			[]int{},
		},
		{
			"prime factor of 2",
			2,
			[]int{2},
		},
		{
			"prime factor of 3",
			3,
			[]int{3},
		},
		{
			"prime factor of 4",
			4,
			[]int{2, 2},
		},
		{
			"prime factor of 5",
			5,
			[]int{5},
		},
		{
			"prime factor of 6",
			6,
			[]int{2, 3},
		},
		{
			"prime factor of 7",
			7,
			[]int{7},
		},
		{
			"prime factor of 8",
			8,
			[]int{2, 2, 2},
		},
		{
			"prime factor of 9",
			9,
			[]int{3, 3},
		},
		{
			"prime factor of a very large number",
			2*2*3*3*3*5*11*17,
			[]int{2, 2, 3, 3, 3, 5, 11, 17},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := PrimeFactorOf(tc.input)
			if !reflect.DeepEqual(output, tc.output) {
				t.Errorf("PrimeFactorOf(%d)=%v, want %v", tc.input, output, tc.output)
			}
		})
	}
}
