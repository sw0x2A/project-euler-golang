// https://projecteuler.net/problem=16

package main

func p16(e int) int {
	sum, al := 0, (e/3)+1
	n := make([]int, al)
	n[al-2], n[al-1] = 2, 0

	for a := 1; a < e; a++ {
		for i := 0; i < al-1; i++ {
			if n[i+1] > 4 {
				n[i] = ((2 * n[i]) % 10) + 1
			} else {
				n[i] = (2 * n[i]) % 10
			}
		}
	}
	for b := 0; b < al-1; b++ {
		sum = sum + n[b]
	}
	return sum
}
