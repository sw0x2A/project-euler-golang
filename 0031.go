// https://projecteuler.net/problem=31

package main

func p31(n int) int {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, n+1)
	ways[0] = 1
	for i := range coins {
		for j := coins[i]; j <= n; j++ {
			ways[j] = ways[j] + ways[j-coins[i]]
		}
	}
	return ways[n]
}
