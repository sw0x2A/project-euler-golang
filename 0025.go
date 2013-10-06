// https://projecteuler.net/problem=25

package main

import (
	"math/big"
)

func p25(n int) int {
	a, b, c, counter := big.NewInt(int64(1)), big.NewInt(int64(2)), big.NewInt(int64(0)), 2
	for len(a.String()) < n {
		c.Set(a)
		a.Add(a, b)
		b.Set(c)
		counter++
	}
	return counter
}
