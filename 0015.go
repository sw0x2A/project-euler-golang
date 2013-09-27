// https://projecteuler.net/problem=15

package main

func p15(x uint64) uint64 {
	var r, n, k, i uint64
	r, n, k = 1, 2*x, x
	for i = 1; i <= k; i++ {
		r *= n + 1 - i
		r /= i
	}
	return r
}
