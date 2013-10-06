// https://projecteuler.net/problem=21

package main

func sumOfDivisors(n int) int {
	sum, p := 1, 2
	for p*p <= n && n > 1 {
		if n%p == 0 {
			j := p * p
			n /= p
			for n%p == 0 {
				j *= p
				n /= p
			}
			sum *= j - 1
			sum /= p - 1
		}
		if p == 2 {
			p = 3
		} else {
			p += 2
		}
	}
	if n > 1 {
		sum *= n + 1
	}
	return sum
}

func sumOfProperDivisors(n int) int {
	return sumOfDivisors(n) - n
}

func p21(limit int) int {
	sum := 0
	for a := 2; a < limit; a++ {
		b := sumOfProperDivisors(a)
		if b > a {
			if sumOfProperDivisors(b) == a {
				sum += a + b
			}
		}
	}
	return sum
}
