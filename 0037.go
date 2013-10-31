// http://projecteuler.net/problem=37

package main

func p37() int {
	sum := 0
	primes := PrimesBelow(750000)
Loop:
	for _, p := range primes {
		if p < 11 {
			continue
		}
		left, right := p, p
		for left > 0 {
			if !IsPrime(left, primes) {
				continue Loop
			}
			left /= 10
		}
		for right > 0 {
			if !IsPrime(right, primes) {
				continue Loop
			}
			switch {
			case right > 100000:
				right %= 100000
			case right > 10000:
				right %= 10000
			case right > 1000:
				right %= 1000
			case right > 100:
				right %= 100
			case right > 10:
				right %= 10
			case right > 1:
				right %= 1
			}
		}
		sum += p
	}
	return sum
}
