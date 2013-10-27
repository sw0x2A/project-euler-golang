// https://projecteuler.net/problem=27

package main

var primes []int

func p27() int {
	primes = PrimesBelow(87400)
	bPos := PrimesBelow(1000)
	max, p := 0, 0
	for a := -999; a < 1000; a += 2 {
		for i := 1; i < len(bPos); i++ {
			for j := 0; j < 2; j++ {
				n := 0
				sign := 1
				if j != 0 {
					sign = -1
				}
				aodd := 0
				if i%2 == 0 {
					aodd = -1
				}
				for isPrime(abs(n*n + (a+aodd)*n + sign*bPos[i])) {
					n++
				}
				if max < n {
					max = n
					p = a * bPos[i]
				}
			}
		}
	}
	return p
}

func isPrime(n int) bool {
	if n%2 == 0 && n != 2 {
		return false
	}
	for _, p := range primes {
		if p == n {
			return true
		}
		if p > n {
			break
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
