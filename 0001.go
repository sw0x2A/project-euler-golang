// If we list all the natural numbers below 10 that are multiples of 3 or
// 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

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

func SumDivisibleBy(target, n int) int {
	p := target / n
	return n * (p * (p + 1)) / 2
}

func p1fast(target int) int {
	return SumDivisibleBy(target, 3) + SumDivisibleBy(target, 5) - SumDivisibleBy(target, 15)
}

func main() {
	fmt.Println(p1fast(999))
}
