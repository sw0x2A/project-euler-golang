// The following iterative sequence is defined for the set of positive integers:
//
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following
// sequence:
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//
// It can be seen that this sequence (starting at 13 and finishing at 1)
// contains 10 terms. Although it has not been proved yet (Collatz Problem),
// it is thought that all starting numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

func p14(n int) int {
	arr := make([]int, n)
	steps, maxSteps, cur, max, notFound := 0, 0, 0, 0, true
	for i := 1; i < len(arr); i++ {
		notFound = true
		cur = i
		steps = 0
		for notFound {
			if cur == 1 {
				notFound = false
				arr[i] = steps
				if steps > maxSteps {
					maxSteps = steps
					max = i
				}
			} else if cur < i {
				notFound = false
				steps = steps + arr[cur]
				arr[i] = steps
				if steps > maxSteps {
					maxSteps = steps
					max = i
				}
			} else {
				if cur%2 == 0 {
					cur /= 2
				} else {
					cur = 3*cur + 1
				}
			}
			steps++
		}
	}
	return max
}
