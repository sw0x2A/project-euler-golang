// https://projecteuler.net/problem=32

package main

func p32() int {
	products := make(map[int]bool)
	sum := 0
	var prod, compiled int

	for m := 2; m < 200; m++ {
		nbegin := 1234
		if m > 9 {
			nbegin = 123
		}
		nend := 10000/m + 1
		for n := nbegin; n < nend; n++ {
			prod = m * n
			compiled = concat(concat(prod, n), m)
			if compiled >= 1E8 && compiled < 1E9 && isPandigital(uint(compiled)) {
				products[prod] = true
			}
		}
	}
	for i := range products {
		sum += i
	}
	return sum
}

func isPandigital(n uint) bool {
	var digits, count, tmp uint
	digits, count = 0, 0
	for n > 0 {
		tmp = digits
		digits = digits | 1<<((n%10)-1)
		if tmp == digits {
			return false
		}
		count++
		n /= 10
	}
	return digits == (1<<count)-1
}

func concat(a, b int) int {
	c := b
	for c > 0 {
		a *= 10
		c /= 10
	}
	return a + b
}
