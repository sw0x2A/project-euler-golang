// https://projecteuler.net/problem=30

package main

func p30(p int) int {
	sum := 0
	for i := 0; i < (p-1)*PowInt(9, p); i++ {
		if equalPows(i, p) {
			sum += i
		}
	}
	return sum
}

func equalPows(n, p int) bool {
	if n < 2 {
		return false
	}
	nn, d, sum := n, 0, 0
	for nn > 0 {
		d = nn % 10
		nn = nn / 10
		sum += PowInt(d, p)
	}
	return n == sum
}

func PowInt(x, y int) int {
	power := x
	for i := 1; i < y; i++ {
		power *= x
	}
	return power
}
