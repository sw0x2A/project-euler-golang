// https://projecteuler.net/problem=4

package main

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = 10*r + n%10
		n /= 10
	}
	return r
}

func isPalindrome(n int) bool {
	return n == reverse(n)
}

func p4() int {
	var lp, a, b, db int
	lp, a = 0, 999
	for a >= 100 {
		if a%11 == 0 {
			b, db = 999, 1
		} else {
			b, db = 990, 11
		}
		for b >= a {
			if a*b <= lp {
				break
			}
			if isPalindrome(a * b) {
				lp = a * b
			}
			b = b - db
		}
		a--
	}
	return lp
}
