// https://projecteuler.net/problem=20

package main

import (
	"math/big"
	"strconv"
)

func factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)
	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}

func p20(n int) int {
	r := big.NewInt(int64(n))
	s, sum := factorial(r).String(), 0
	for i := 0; i < len(s); i++ {
		d, _ := strconv.Atoi(string(s[i]))
		sum += d
	}
	return sum
}
