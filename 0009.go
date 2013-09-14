package main

import "math"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func p9(s int) int {
	s2 := s / 2
	mlimit := int(math.Sqrt(float64(s2))) - 1
	var a, b, c, k int
loop:
	for m := 2; m <= mlimit; m++ {
		if s2%m == 0 {
			sm := s2 / m
			for sm%2 == 0 {
				sm /= 2
			}
			if m%2 == 1 {
				k = m + 2
			} else {
				k = m + 1
			}
			for k < 2*m && k <= sm {
				if sm%k == 0 && gcd(k, m) == 1 {
					d := s2 / (k * m)
					n := k - m
					a = d * (m*m - n*n)
					b = 2 * d * m * n
					c = d * (m*m + n*n)
					break loop
				}
				k += 2
			}
		}
	}
	return a * b * c
}
