// https://projecteuler.net/problem=1

package main

// Easy solution; not used
func p1easy(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func p1fast(target int) int {
	// sdb: SumDivisibleBy
	sdb := func(n int) int {
		p := target / n
		return n * (p * (p + 1)) / 2
	}
	return sdb(3) + sdb(5) - sdb(15)
}
