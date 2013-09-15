// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import "math"

func p10(limit int) int {
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
	sum := 2
	for i := 1; i <= sievebound; i++ {
		if !sieve[i] {
			sum += 2*i + 1
		}
	}
	return sum
}
