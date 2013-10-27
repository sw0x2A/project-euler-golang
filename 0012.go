// https://projecteuler.net/problem=12

package main

// Requires PrimesBelow from utils.go
func p12() int {
	n, dn, cnt := 3, 2, 0
	primearray := PrimesBelow(1000000)
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
