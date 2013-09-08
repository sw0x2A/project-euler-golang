// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

func p3(n int64) int {
	var f, lf, mf int64
	if n%2 == 0 {
		lf = 2
		n /= 2
		for n%2 == 0 {
			n /= 2
		}
	} else {
		lf = 1
	}
	f = 3
	mf = int64(math.Sqrt(float64(n)))
	for n > 1 && f <= mf {
		if n%f == 0 {
			n /= f
			lf = f
			for n%f == 0 {
				n /= f
			}
			mf = int64(math.Sqrt(float64(n)))
		}
		f += 2
	}
	if n == 1 {
		return int(lf)
	}
	return int(n)
}

func main() {
	fmt.Println(p3(600851475143))
}
