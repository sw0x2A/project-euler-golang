// https://projecteuler.net/problem=2

package main

func p2(n int) int {
	sum, a, b, c := 0, 1, 1, 2 // c=a+b
	for c < n {
		sum += c
		a = b + c
		b = c + a
		c = a + b
	}
	return sum
}
