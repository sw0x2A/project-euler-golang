package main

import "strconv"

func p8(s string) int {
	var n [1000]int
	for j := 0; j < len(s); j++ {
		n[j], _ = strconv.Atoi(string(s[j]))
	}
	var largest, t int
	for i := 0; i <= len(s)-5; i++ {
		t = n[i] * n[i+1] * n[i+2] * n[i+3] * n[i+4]
		if largest < t {
			largest = t
		}
	}
	return largest
}
