// https://projecteuler.net/problem=28

package main

// Numbers which are on the diagonals in Ulam's spiral
// http://oeis.org/A200975
func p28(n int) int {
	n = n / 2
	return 8*n*(n+1)*(2*n+1)/3 + 2*n*(n+1) + 4*n + 1
}
