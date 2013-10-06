// https://projecteuler.net/problem=24

package main

func p24(n, i int) int {
	j, k := 0, 0
	fact := make([]int, n)
	perm := make([]int, n)
	fact[k] = 1
	k++
	i--
	for k < n {
		fact[k] = fact[k-1] * k
		k++
	}
	for k = 0; k < n; k++ {
		perm[k] = i / fact[n-1-k]
		i = i % fact[n-1-k]
	}
	for k = n - 1; k > 0; k-- {
		for j = k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}
	result := 0
	for k = 0; k < n; k++ {
		result = (result + perm[k]) * 10
	}
	return result / 10
}
