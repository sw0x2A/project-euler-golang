// https://projecteuler.net/problem=6

package main

func p6(n uint) uint {
	sum := n * (n + 1) / 2
	sum_sq := (2*n + 1) * (n + 1) * n / 6
	return sum*sum - sum_sq
}
