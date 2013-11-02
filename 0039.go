// http://projecteuler.net/problem=39

package main

import "math"

func p39() int {
	pmax, tmax, k := 0, 0, 0
	for s := 1; s <= 1000; s++ {
		t := 0
		mlimit := int(math.Sqrt(float64(s) / 2))
		for m := 2; m <= mlimit; m++ {
			if (s/2)%m == 0 {
				if m%2 == 0 {
					k = m + 1
				} else {
					k = m + 2
				}
				for k < 2*m && k <= s/(2*m) {
					if s/(2*m)%k == 0 && Gcd(k, m) == 1 {
						t++
					}
					k += 2
				}
			}
		}
		if t > tmax {
			tmax = t
			pmax = s
		}
	}
	return pmax
}
