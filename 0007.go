// https://projecteuler.net/problem=7

package main

func p7(n int) int {
	s := make([]bool, 105000)
	s[0], s[1] = true, true
	i, p := 2, 3
	var k int
	for {
		for k = 2 * p; k < len(s); k += p {
			s[k] = true
		}
		for k = p + 2; k < len(s) && s[k]; k += 2 {
		}
		if k < len(s) {
			p = k
			i++
			if i == n {
				break
			}
		} else {
			break
		}
	}
	return p
}
