// https://projecteuler.net/problem=27

package main

import (
	"math"
)

// Sieve of Eratosthenes
func PrimesBelow(limit int) []int {
	switch {
	case limit < 2:
		return []int{}
	case limit == 2:
		return []int{2}
	}
	sievebound := (limit - 1) / 2
	sieve := make([]bool, sievebound+1)
	crosslimit := int(math.Sqrt(float64(limit))-1) / 2
	for i := 1; i <= crosslimit; i++ {
		if !sieve[i] {
			for j := 2 * i * (i + 1); j <= sievebound; j += 2*i + 1 {
				sieve[j] = true
			}
		}
	}
	plimit := int(1.3*float64(limit)) / int(math.Log(float64(limit)))
	primes := make([]int, plimit)
	p := 1
	primes[0] = 2
	for i := 1; i <= sievebound; i++ {
		if !sieve[i] {
			primes[p] = 2*i + 1
			p++
			if p >= plimit {
				break
			}
		}
	}
	last := len(primes) - 1
	for i := last; i > 0; i-- {
		if primes[i] != 0 {
			break
		}
		last = i
	}
	return primes[0:last]
}
