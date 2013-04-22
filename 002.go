package main

import "strconv"

func Problem002() string {
	a := 1
	b := 2
	sum := 0

	limit := 4000000
	for a < limit || b < limit {
		if b%2 == 0 {
			sum += b
		}
		b = a + b
		a = b - a
	}

	return strconv.Itoa(sum)
}
