// https://projecteuler.net/problem=12

package main

import "math"

func p12() int {
	n, dn, cnt := 3, 2, 0
	primearray := sieveOfAtkin()
	for cnt <= 500 {
		n++
		n1 := n
		if n1%2 == 0 {
			n1 /= 2
		}
		dn1 := 1
		for i := 0; i < len(primearray); i++ {
			if primearray[i]*primearray[i] > n1 {
				dn1 *= 2
				break
			}
			exponent := 1
			for n1%primearray[i] == 0 {
				exponent++
				n1 /= primearray[i]
			}
			if exponent > 1 {
				dn1 *= exponent
			}
			if n1 == 1 {
				break
			}
		}
		cnt = dn * dn1
		dn = dn1
	}
	return n * (n - 1) / 2
}

func sieveOfAtkin() []int {
	limit := 1000000
	limitSqrt := int(math.Sqrt(float64(limit)))
	sieve := make([]bool, limit+1)
	sieve[0], sieve[1], sieve[2], sieve[3] = false, false, true, true
	for x := 1; x <= limitSqrt; x++ {
		for y := 1; y <= limitSqrt; y++ {
			n := 4*x*x + y*y
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				sieve[n] = !sieve[n]
			}

			n = 3*x*x + y*y
			if n <= limit && n%12 == 7 {
				sieve[n] = !sieve[n]
			}

			n = 3*x*x - y*y
			if x > y && n <= limit && n%12 == 11 {
				sieve[n] = !sieve[n]
			}
		}
	}
	for n := 5; n <= limitSqrt; n++ {
		if sieve[n] {
			x := n * n
			for i := x; i <= limit; i += x {
				sieve[i] = false
			}
		}
	}
	np := 65500
	primes := make([]int, np)
	c := 0
	for i := 0; i <= limit; i++ {
		if sieve[i] {
			primes[c] = i
			c++
			if c >= np {
				break
			}
		}
	}
	return primes
}
