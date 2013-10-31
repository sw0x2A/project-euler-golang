// https://projecteuler.net/problem=27

package main

func p27() int {
	primes := PrimesBelow(87400)
	bPos := primes[0:168]
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
				for IsPrime(abs(n*n+(a+aodd)*n+sign*bPos[i]), primes) {
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

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
