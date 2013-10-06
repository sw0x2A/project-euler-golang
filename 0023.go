// https://projecteuler.net/problem=23

package main

// Requires functions sumOfDivisors and sumOfProperDivisors from solution of problem 21
func isAbundant(n int) bool {
	if n < sumOfProperDivisors(n) {
		return true
	}
	return false
}

func p23() int {
	limit := 28124
	abundants := make([]int, 6965)
	ai := 0
	for i := 2; i < limit; i++ {
		if isAbundant(i) {
			abundants[ai] = i
			ai++
		}
	}
	sumOfAbundants := make([]bool, limit)
	for _, v := range abundants {
		for _, w := range abundants {
			if sum := v + w; sum < limit {
				sumOfAbundants[sum] = true
			}
		}
	}
	sum := 0
	for i := range sumOfAbundants {
		if !sumOfAbundants[i] {
			sum += i
		}
	}
	return sum
}
