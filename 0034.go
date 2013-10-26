// https://projecteuler.net/problem=34

package main

func p34() int {
	facts := make([]int, 10)
	facts[0], facts[1] = 1, 1
	for i := 2; i < 10; i++ {
		facts[i] = facts[i-1] * i
	}
	sum := 0
	// 9!*7 = 2540160
	for i := 10; i < 2540161; i++ {
		sumOfFacts, number := 0, i
		for number > 0 {
			d := number % 10
			number /= 10
			sumOfFacts += facts[d]
		}
		if sumOfFacts == i {
			sum += i
		}
	}
	return sum
}
