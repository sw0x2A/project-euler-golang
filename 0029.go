// https://projecteuler.net/problem=29

package main

import "math/big"

func p29(n int) int {
	set := make(map[string]bool)
	for a := 2; a <= n; a++ {
		for b := 2; b <= n; b++ {
			ab := big.NewInt(int64(a))
			bb := big.NewInt(int64(b))
			set[ab.Exp(ab, bb, nil).String()] = true
		}
	}
	return len(set)
}
